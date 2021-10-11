package keycloakrealmrolebatch

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1alpha1"
	"github.com/epam/edp-keycloak-operator/pkg/client/keycloak/mock"
	"github.com/epam/edp-keycloak-operator/pkg/controller/helper"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	k8sCLient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func TestReconcileKeycloakRealmRoleBatch_ReconcileDelete(t *testing.T) {
	scheme := runtime.NewScheme()
	utilruntime.Must(v1alpha1.AddToScheme(scheme))
	utilruntime.Must(corev1.AddToScheme(scheme))

	ns := "security"
	keycloak := v1alpha1.Keycloak{
		ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: ns}, Status: v1alpha1.KeycloakStatus{Connected: true}}
	realm := v1alpha1.KeycloakRealm{ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: ns,
		OwnerReferences: []metav1.OwnerReference{{Name: "test", Kind: "Keycloak"}}},
		Spec: v1alpha1.KeycloakRealmSpec{RealmName: "test"}}
	secret := corev1.Secret{ObjectMeta: metav1.ObjectMeta{Name: "keycloak-secret", Namespace: ns},
		Data: map[string][]byte{"username": []byte("user"), "password": []byte("pass")}}
	now := metav1.Time{Time: time.Now()}
	batch := v1alpha1.KeycloakRealmRoleBatch{ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: ns,
		DeletionTimestamp: &now,
		OwnerReferences:   []metav1.OwnerReference{{Name: "test", Kind: "KeycloakRealm"}}},
		Spec: v1alpha1.KeycloakRealmRoleBatchSpec{Realm: "test", Roles: []v1alpha1.BatchRole{
			{Name: "sub-role1"},
			{Name: "sub-role2", IsDefault: true},
		}},
		Status: v1alpha1.KeycloakRealmRoleBatchStatus{}}

	client := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(&batch, &realm, &keycloak, &secret).Build()

	rkr := ReconcileKeycloakRealmRoleBatch{scheme: scheme, client: client, helper: helper.MakeHelper(client, scheme, nil),
		log: &mock.Logger{}}

	if _, err := rkr.Reconcile(context.Background(), reconcile.Request{
		NamespacedName: types.NamespacedName{
			Name:      "test",
			Namespace: ns,
		},
	}); err != nil {
		t.Fatal(err)
	}

	var checkList v1alpha1.KeycloakRealmRoleList
	if err := client.List(context.Background(), &checkList, &k8sCLient.ListOptions{}); err != nil {
		t.Fatal(err)
	}

	if len(checkList.Items) > 0 {
		t.Fatal("batch roles is not deleted")
	}

}

func TestReconcileKeycloakRealmRoleBatch_Reconcile(t *testing.T) {
	scheme := runtime.NewScheme()
	utilruntime.Must(v1alpha1.AddToScheme(scheme))
	utilruntime.Must(corev1.AddToScheme(scheme))

	ns := "security"
	keycloak := v1alpha1.Keycloak{
		ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: ns}, Status: v1alpha1.KeycloakStatus{Connected: true}}
	realm := v1alpha1.KeycloakRealm{ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: ns,
		OwnerReferences: []metav1.OwnerReference{{Name: "test", Kind: "Keycloak"}}},
		Spec: v1alpha1.KeycloakRealmSpec{RealmName: "test"}}
	secret := corev1.Secret{ObjectMeta: metav1.ObjectMeta{Name: "keycloak-secret", Namespace: ns},
		Data: map[string][]byte{"username": []byte("user"), "password": []byte("pass")}}
	batch := v1alpha1.KeycloakRealmRoleBatch{ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: ns,
		OwnerReferences: []metav1.OwnerReference{{Name: "test", Kind: "KeycloakRealm"}}},
		Spec: v1alpha1.KeycloakRealmRoleBatchSpec{Realm: "test", Roles: []v1alpha1.BatchRole{
			{Name: "sub-role1"},
			{Name: "sub-role2", IsDefault: true},
		}},
		Status: v1alpha1.KeycloakRealmRoleBatchStatus{}}

	role := v1alpha1.KeycloakRealmRole{ObjectMeta: metav1.ObjectMeta{Name: "test2", Namespace: ns,
		OwnerReferences: []metav1.OwnerReference{
			{Name: "test", Kind: "KeycloakRealm"},
			{Name: "test", Kind: batch.Kind},
		}},
		Spec:   v1alpha1.KeycloakRealmRoleSpec{Name: "test"},
		Status: v1alpha1.KeycloakRealmRoleStatus{Value: ""},
	}

	client := fake.NewClientBuilder().WithScheme(scheme).
		WithRuntimeObjects(&batch, &realm, &keycloak, &secret, &role).Build()

	rkr := ReconcileKeycloakRealmRoleBatch{
		scheme: scheme,
		client: client,
		helper: helper.MakeHelper(client, scheme, nil),
		log:    &mock.Logger{}}

	if _, err := rkr.Reconcile(context.TODO(), reconcile.Request{
		NamespacedName: types.NamespacedName{
			Name:      "test",
			Namespace: ns,
		},
	}); err != nil {
		t.Fatal(err)
	}

	var checkBatch v1alpha1.KeycloakRealmRoleBatch
	if err := client.Get(context.Background(), types.NamespacedName{
		Namespace: ns,
		Name:      "test",
	}, &checkBatch); err != nil {
		t.Fatal(err)
	}

	if checkBatch.Status.Value != helper.StatusOK {
		t.Log(checkBatch.Status.Value)
		t.Fatal("batch status not updated on success")
	}

	var roles v1alpha1.KeycloakRealmRoleList
	if err := client.List(context.Background(), &roles, &k8sCLient.ListOptions{}); err != nil {
		t.Fatal(err)
	}

	var checkRole v1alpha1.KeycloakRealmRole
	if err := client.Get(context.Background(), types.NamespacedName{Namespace: ns,
		Name: fmt.Sprintf("%s-sub-role2", batch.Name)}, &checkRole); err != nil {
		t.Fatal(err)
	}

	if !checkRole.Spec.IsDefault {
		t.Fatal("sub-role2 is not default")
	}

	checkBatch.Spec.Roles = checkBatch.Spec.Roles[1:]
	if err := client.Update(context.Background(), &checkBatch); err != nil {
		t.Fatal(err)
	}

	if _, err := rkr.Reconcile(context.TODO(), reconcile.Request{
		NamespacedName: types.NamespacedName{
			Name:      "test",
			Namespace: ns,
		},
	}); err != nil {
		t.Fatal(err)
	}

	if err := client.Get(context.Background(), types.NamespacedName{Namespace: ns,
		Name: fmt.Sprintf("%s-sub-role1", batch.Name)}, &checkRole); err == nil {
		t.Fatal("sub role is not marked for deletion")
	}
}

func TestReconcileKeycloakRealmRoleBatch_ReconcileFailure(t *testing.T) {
	scheme := runtime.NewScheme()
	utilruntime.Must(v1alpha1.AddToScheme(scheme))
	utilruntime.Must(corev1.AddToScheme(scheme))

	ns := "security"
	keycloak := v1alpha1.Keycloak{
		ObjectMeta: metav1.ObjectMeta{Name: "keycloak1", Namespace: ns}, Status: v1alpha1.KeycloakStatus{Connected: true}}
	realm := v1alpha1.KeycloakRealm{ObjectMeta: metav1.ObjectMeta{Name: "realm1", Namespace: ns,
		OwnerReferences: []metav1.OwnerReference{{Name: "keycloak1", Kind: "Keycloak"}}},
		Spec: v1alpha1.KeycloakRealmSpec{RealmName: "ns-realm1"}}
	secret := corev1.Secret{ObjectMeta: metav1.ObjectMeta{Name: "keycloak-secret", Namespace: ns},
		Data: map[string][]byte{"username": []byte("user"), "password": []byte("pass")}}
	now := metav1.Time{Time: time.Now()}
	batch := v1alpha1.KeycloakRealmRoleBatch{ObjectMeta: metav1.ObjectMeta{Name: "batch1", Namespace: ns,
		DeletionTimestamp: &now},
		Spec: v1alpha1.KeycloakRealmRoleBatchSpec{Realm: "realm1", Roles: []v1alpha1.BatchRole{
			{Name: "role1"},
			{Name: "role2"},
		}},
		Status: v1alpha1.KeycloakRealmRoleBatchStatus{}}

	role := v1alpha1.KeycloakRealmRole{ObjectMeta: metav1.ObjectMeta{Name: "batch1-role2", Namespace: ns},
		Spec:   v1alpha1.KeycloakRealmRoleSpec{Name: "batch1-role2", Realm: "realm1"},
		Status: v1alpha1.KeycloakRealmRoleStatus{Value: ""},
	}

	client := fake.NewClientBuilder().WithScheme(scheme).
		WithRuntimeObjects(&batch, &realm, &keycloak, &secret, &role).Build()

	logger := mock.Logger{}
	rkr := ReconcileKeycloakRealmRoleBatch{
		scheme: scheme,
		client: client,
		helper: helper.MakeHelper(client, scheme, &logger),
		log:    &logger,
	}

	_, err := rkr.Reconcile(context.TODO(), reconcile.Request{
		NamespacedName: types.NamespacedName{
			Name:      "batch1",
			Namespace: ns,
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	logErr := logger.LastError()

	if logErr == nil {
		t.Fatal("no error on fatal")
	}

	if errors.Cause(logErr).Error() != "one of batch role already exists" {
		t.Fatal("wrong error returned")
	}

	var checkBatch v1alpha1.KeycloakRealmRoleBatch
	if err := client.Get(context.Background(), types.NamespacedName{
		Namespace: ns,
		Name:      "batch1",
	}, &checkBatch); err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(checkBatch.Status.Value, "one of batch role already exists") {
		t.Log(checkBatch.Status.Value)
		t.Fatal("batch status not updated on failure")
	}
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"k8s.io/apimachinery/pkg/types"
	"log"
	"math/rand"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"time"

	//apierrs "k8s.io/apimachinery/pkg/api/errors"
	//"k8s.io/apimachinery/pkg/types"

	hydrav1alpha1 "github.com/ory/hydra-maester/api/v1alpha1"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	// +kubebuilder:scaffold:imports
)

var (
	scheme   = runtime.NewScheme()
)

func init() {

	apiv1.AddToScheme(scheme)
	hydrav1alpha1.AddToScheme(scheme)
	// +kubebuilder:scaffold:scheme
}

func main() {

	ctx := context.Background()

	log.Println("starting sync...")

	c, err := client.New(ctrl.GetConfigOrDie(), client.Options{Scheme:scheme})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var oauth2list hydrav1alpha1.OAuth2ClientList
	if err := c.List(ctx, &oauth2list, client.InNamespace("default")); err != nil {
		log.Fatal("failed to list clients")
	}

	str:=fmt.Sprintf(`{
			"spec" : {
				"metadata": {
					"sync": "%s"
				}
			}
		}`, generateRandomString(4))

	data:=json.RawMessage(str)

	for _, oauth2client := range oauth2list.Items {
		log.Println(oauth2client.Name)
		if err := c.Patch(ctx, &oauth2client, client.ConstantPatch(types.MergePatchType, data)); err != nil {
			fmt.Println(err)
		}
	}
}

func generateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())

	letterRunes := []rune("abcdefghijklmnopqrstuvwxyz")

	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
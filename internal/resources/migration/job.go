package migration

import (
	"fmt"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/ZupIT/horusec-operator/api/v2alpha1"
)

func NewJob(resource *v2alpha1.HorusecPlatform) batchv1.Job {
	var terminationPeriod int64 = 30
	global := resource.Spec.Global
	return batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: fmt.Sprintf("%s-platform-migration-", resource.GetName()),
			Namespace:    resource.GetNamespace(),
			Labels:       resource.GetDefaultLabel(),
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					RestartPolicy:                 corev1.RestartPolicyOnFailure,
					TerminationGracePeriodSeconds: &terminationPeriod,
					Containers: []corev1.Container{
						{
							Name:            "horusec-database-migration",
							Image:           resource.GetDatabaseMigrationImage(),
							ImagePullPolicy: corev1.PullIfNotPresent,
							Command:         []string{"migrate.sh"},
							Env: []corev1.EnvVar{
								resource.NewEnvFromSecret("HORUSEC_DATABASE_USERNAME", global.Database.User.KeyRef),
								resource.NewEnvFromSecret("HORUSEC_DATABASE_PASSWORD", global.Database.Password.KeyRef),
								{Name: "MIGRATION_NAME", Value: "platform"},
								{Name: "HORUSEC_DATABASE_SQL_URI", Value: resource.GetGlobalDatabaseURI()},
							},
						},
					},
				},
			},
		},
	}
}

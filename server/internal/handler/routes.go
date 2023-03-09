// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	configmaps "server/internal/handler/configmaps"
	daemonSets "server/internal/handler/daemonSets"
	deployments "server/internal/handler/deployments"
	ingresses "server/internal/handler/ingresses"
	namespaces "server/internal/handler/namespaces"
	nodes "server/internal/handler/nodes"
	pods "server/internal/handler/pods"
	pvcs "server/internal/handler/pvcs"
	pvs "server/internal/handler/pvs"
	secrets "server/internal/handler/secrets"
	services "server/internal/handler/services"
	statefulSets "server/internal/handler/statefulSets"
	"server/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/pods",
				Handler: pods.GetPodListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/pod/detail",
				Handler: pods.GetPodDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/pod/del",
				Handler: pods.DeletePodHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/pod/update",
				Handler: pods.UpdatePodHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/pod/container",
				Handler: pods.GetPodContainersHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/pod/log",
				Handler: pods.GetPodLogHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/k8s"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/service/create",
				Handler: services.CreateServiceHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/services",
				Handler: services.GetServiceListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/service/detail",
				Handler: services.GetServiceDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/service/del",
				Handler: services.DeleteServiceHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/service/update",
				Handler: services.UpdateServiceHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/k8s"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/deployments",
				Handler: deployments.GetDeploymentListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/deployment/detail",
				Handler: deployments.GetDeploymentDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/deployment/del",
				Handler: deployments.DeleteDeploymentHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/deployment/update",
				Handler: deployments.UpdateDeploymentHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/deployment/scale",
				Handler: deployments.ScaleDeploymentHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/deployment/restart",
				Handler: deployments.RestartDeploymentHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/deployment/create",
				Handler: deployments.CreateDeploymentHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/k8s"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/daemonsets",
				Handler: daemonSets.GetDaemonSetListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/daemonset/detail",
				Handler: daemonSets.GetDaemonSetDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/daemonset/delete",
				Handler: daemonSets.DeleteDaemonSetHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/daemonset/update",
				Handler: daemonSets.UpdateDaemonSetHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/k8s"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/statefulsets",
				Handler: statefulSets.GetStatefulSetListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/statefulset/detail",
				Handler: statefulSets.GetStatefulSetDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/statefulset/delete",
				Handler: statefulSets.DeleteStatefulSetHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/statefulset/update",
				Handler: statefulSets.UpdateStatefulSetHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/k8s"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/ingress/create",
				Handler: ingresses.CreateIngressHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/ingresses",
				Handler: ingresses.GetIngressListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/ingress/detail",
				Handler: ingresses.GetIngressDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/ingress/delete",
				Handler: ingresses.DeleteIngressHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/ingress/update",
				Handler: ingresses.UpdateIngressHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/k8s"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/nodes",
				Handler: nodes.GetNodesHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/node/detail",
				Handler: nodes.GetNodeDetailHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/k8s"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/namespaces",
				Handler: namespaces.GetNamespacesHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/namespace/detail",
				Handler: namespaces.GetNamespaceDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/namespace/del",
				Handler: namespaces.DeleteNamespaceHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/k8s"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/pvs",
				Handler: pvs.GetPVsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/pv/detail",
				Handler: pvs.GetPVDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/pv/del",
				Handler: pvs.DeletePVHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/k8s"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/configmaps",
				Handler: configmaps.GetConfigmapsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/configmap/detail",
				Handler: configmaps.GetConfigmapDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/configmap/del",
				Handler: configmaps.DeleteConfigmapHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/configmap/update",
				Handler: configmaps.UpdateConfigmapHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/k8s"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/secrets",
				Handler: secrets.GetSecretsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/secret/detail",
				Handler: secrets.GetSecretDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/secret/del",
				Handler: secrets.DeleteSecretHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/secret/update",
				Handler: secrets.UpdateSecretHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/k8s"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/pvcs",
				Handler: pvcs.GetPVCsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/pvc/detail",
				Handler: pvcs.GetPVCDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/pvc/del",
				Handler: pvcs.DeletePVCHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/pvc/update",
				Handler: pvcs.UpdatePVCHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/k8s"),
	)
}

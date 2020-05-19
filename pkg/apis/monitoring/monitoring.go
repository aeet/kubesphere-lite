/*
 * @File: monitoring
 * @Date: 2020/5/18 3:14 下午
 * @Author: ferried
 * @Email: harlancui@outlook.com
 * @Software: GoLand
 * @Desc: null
 * @License: null
 */
package monitoring

import (
	"github.com/gin-gonic/gin"
	"my-kubesphere/pkg/apiserver"
	"my-kubesphere/pkg/metrics"
	"net/url"
	"strconv"
	"strings"
)

func MonitorCluster(ctx *gin.Context) {
	r := ParseRequestParams(ctx)
	var res *metrics.Response
	res = metrics.GetClusterMetrics(r)
	ctx.JSON(200, res)
}

func ParseRequestParams(c *gin.Context) metrics.RequestParams {
	var requestParams metrics.RequestParams

	queryTime := strings.Trim(c.Query("time"), " ")
	start := strings.Trim(c.Query("start"), " ")
	end := strings.Trim(c.Query("end"), " ")
	step := strings.Trim(c.Query("step"), " ")
	sortMetric := strings.Trim(c.Query("sort_metric"), " ")
	sortType := strings.Trim(c.Query("sort_type"), " ")
	pageNum := strings.Trim(c.Query("page"), " ")
	limitNum := strings.Trim(c.Query("limit"), " ")
	tp := strings.Trim(c.Query("type"), " ")
	metricsFilter := strings.Trim(c.Query("metrics_filter"), " ")
	resourcesFilter := strings.Trim(c.Query("resources_filter"), " ")
	nodeName := strings.Trim(c.Param("node"), " ")
	workspaceName := strings.Trim(c.Param("workspace"), " ")
	namespaceName := strings.Trim(c.Param("namespace"), " ")
	workloadKind := strings.Trim(c.Param("kind"), " ")
	workloadName := strings.Trim(c.Param("workload"), " ")
	podName := strings.Trim(c.Param("pod"), " ")
	containerName := strings.Trim(c.Param("container"), " ")
	pvcName := strings.Trim(c.Param("pvc"), " ")
	storageClassName := strings.Trim(c.Param("storageclass"), " ")
	componentName := strings.Trim(c.Param("component"), " ")

	requestParams = metrics.RequestParams{
		SortMetric:       sortMetric,
		SortType:         sortType,
		PageNum:          pageNum,
		LimitNum:         limitNum,
		Type:             tp,
		MetricsFilter:    metricsFilter,
		ResourcesFilter:  resourcesFilter,
		NodeName:         nodeName,
		WorkspaceName:    workspaceName,
		NamespaceName:    namespaceName,
		WorkloadKind:     workloadKind,
		WorkloadName:     workloadName,
		PodName:          podName,
		ContainerName:    containerName,
		PVCName:          pvcName,
		StorageClassName: storageClassName,
		ComponentName:    componentName,
	}

	if metricsFilter == "" {
		requestParams.MetricsFilter = ".*"
	}
	if resourcesFilter == "" {
		requestParams.ResourcesFilter = ".*"
	}

	v := url.Values{}

	if start != "" && end != "" { // range query

		// metrics from a deleted namespace should be hidden
		// therefore, for range query, if range query start time is less than the namespace creation time, set it to creation time
		// it is the same with query at a fixed time point
		if namespaceName != "" {
			nsLister := apiserver.Helper.InformerFactory.KubernetesSharedInformerFactory().Core().V1().Namespaces().Lister()
			ns, err := nsLister.Get(namespaceName)
			if err == nil {
				creationTime := ns.CreationTimestamp.Time.Unix()
				queryStart, err := strconv.ParseInt(start, 10, 64)
				if err == nil && queryStart < creationTime {
					start = strconv.FormatInt(creationTime, 10)
				}
			}
		}

		v.Set("start", start)
		v.Set("end", end)

		if step == "" {
			v.Set("step", metrics.DefaultQueryStep)
		} else {
			v.Set("step", step)
		}
		requestParams.QueryParams = v
		requestParams.QueryType = metrics.RangeQuery

		return requestParams
	} else if queryTime != "" { // query
		v.Set("time", queryTime)
	}

	requestParams.QueryParams = v
	requestParams.QueryType = metrics.Query
	return requestParams
}

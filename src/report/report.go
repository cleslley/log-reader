package report

import (
	"encoding/csv"
	"fmt"
	"log-reader/src/model"
	"os"
)

func CreateConsumerReport(logs []*model.Log) error {
	outputName := "ReportByConsumer.csv"
	idKind := "consumer"

	ids := getIdLogs(idKind, logs)
	reportData := createReportData(idKind, ids, logs)

	err := createCsvFile(outputName, reportData)
	if err != nil {
		return err
	}

	return nil
}

func CreateServiceReport(logs []*model.Log) error {
	outputName := "ReportByService.csv"
	idKind := "service"

	ids := getIdLogs(idKind, logs)
	reportData := createReportData(idKind, ids, logs)

	err := createCsvFile(outputName, reportData)
	if err != nil {
		return err
	}

	return nil
}

func CreateServiceTimeReport(logs []*model.Log) error {
	outputName := "ReportByServiceTime.csv"
	idKind := "service"

	ids := getIdLogs(idKind, logs)
	reportData := createServiceTimeReportData(ids, logs)

	err := createCsvFile(outputName, reportData)
	if err != nil {
		return err
	}

	return nil
}

func getConsumerReportHeader() []string {
	header := []string{
		"ConsumerID",
		"RequestURI",
		"RequestURL",
		"RequestSize",
		"RequestQueryString",
		"RequestHeaderAccept",
		"RequestHeaderHost",
		"RequestHeaderUserAgent",
		"UpStreamURI",
		"ResponseStatus",
		"ResponseSize",
		"ResponseHeaderContentLength",
		"ResponseHeaderVia",
		"ResponseHeaderConnection",
		"ResponseHeaderAcessControlCred",
		"ResponseHeaderContentyType",
		"ResponseHeaderServer",
		"ResponseHeaderAcessControlOrigin",
		"RouteCreatedAt",
		"RouteHosts",
		"RouteID",
		"RouteMethods",
		"RoutePaths",
		"RoutePreserveHost",
		"RouteProtocols",
		"RouteRegexPriority",
		"RouteServiceID",
		"RouteStripPath",
		"RouteUpdatedAt",
		"ServiceConnectTimeout",
		"ServiceCreatedAt",
		"ServiceHost",
		"ServiceID",
		"ServiceName",
		"ServicePath",
		"ServicePort",
		"ServiceProtocol",
		"ServiceReadTimeout",
		"ServiceRetries",
		"ServiceUpdatedAt",
		"ServiceWriteTimeout",
		"LatenciesProxy",
		"LatenciesKong",
		"LatenciesRequest",
		"ClientIp",
		"StartedAt",
	}

	return header
}

func getServiceReportHeader() []string {
	header := []string{
		"ServiceID",
		"RequestURI",
		"RequestURL",
		"RequestSize",
		"RequestQueryString",
		"RequestHeaderAccept",
		"RequestHeaderHost",
		"RequestHeaderUserAgent",
		"UpStreamURI",
		"ResponseStatus",
		"ResponseSize",
		"ResponseHeaderContentLength",
		"ResponseHeaderVia",
		"ResponseHeaderConnection",
		"ResponseHeaderAcessControlCred",
		"ResponseHeaderContentyType",
		"ResponseHeaderServer",
		"ResponseHeaderAcessControlOrigin",
		"ConsumerID",
		"RouteCreatedAt",
		"RouteHosts",
		"RouteID",
		"RouteMethods",
		"RoutePaths",
		"RoutePreserveHost",
		"RouteProtocols",
		"RouteRegexPriority",
		"RouteServiceID",
		"RouteStripPath",
		"RouteUpdatedAt",
		"ServiceConnectTimeout",
		"ServiceCreatedAt",
		"ServiceHost",
		"ServiceName",
		"ServicePath",
		"ServicePort",
		"ServiceProtocol",
		"ServiceReadTimeout",
		"ServiceRetries",
		"ServiceUpdatedAt",
		"ServiceWriteTimeout",
		"LatenciesProxy",
		"LatenciesKong",
		"LatenciesRequest",
		"ClientIp",
		"StartedAt",
	}

	return header
}

func getServiceTimeReportHeader() []string {
	header := []string{
		"ServiceID",
		"Requests",
		"Average Proxy",
		"Average Gateway",
		"Average Request",
		"Total Average",
	}

	return header
}

func getRowConvertedConsumerReport(log *model.Log) []string {
	row := []string{
		log.AuthenticatedEntity.ConsumerID.Uuid,
		log.Request.Uri,
		log.Request.Url,
		fmt.Sprint(log.Request.Size),
		fmt.Sprint(log.Request.QueryString),
		log.Request.Headers.Accept,
		log.Request.Headers.Host,
		log.Request.Headers.UserAgent,
		log.UpStreamUri,
		fmt.Sprint(log.Response.Status),
		fmt.Sprint(log.Response.Size),
		log.Response.Headers.ContentLength,
		log.Response.Headers.Via,
		log.Response.Headers.Connection,
		log.Response.Headers.AccessControlAllowCredentials,
		log.Response.Headers.ContentType,
		log.Response.Headers.Server,
		log.Response.Headers.AccessControlAllowOrigin,
		fmt.Sprint(log.Route.CreatedAt),
		log.Route.Hosts,
		log.Route.ID,
		fmt.Sprint(log.Route.Methods),
		fmt.Sprint(log.Route.Paths),
		fmt.Sprint(log.Route.PreserveHost),
		fmt.Sprint(log.Route.Protocols),
		fmt.Sprint(log.Route.RegexPriority),
		fmt.Sprint(log.Route.Service.ID),
		fmt.Sprint(log.Route.StripPath),
		fmt.Sprint(log.Route.UpdatedAt),
		fmt.Sprint(log.Service.ConnectTimeout),
		fmt.Sprint(log.Service.CreatedAt),
		fmt.Sprint(log.Service.Host),
		log.Service.ID,
		log.Service.Name,
		log.Service.Path,
		fmt.Sprint(log.Service.Port),
		log.Service.Protocol,
		fmt.Sprint(log.Service.ReadTimeout),
		fmt.Sprint(log.Service.Retries),
		fmt.Sprint(log.Service.UpdatedAt),
		fmt.Sprint(log.Service.WriteTimeout),
		fmt.Sprint(log.Latencies.Proxy),
		fmt.Sprint(log.Latencies.Kong),
		fmt.Sprint(log.Latencies.Request),
		log.ClientIp,
		fmt.Sprint(log.StartedAt),
	}

	return row
}

func getRowConvertedServiceReport(log *model.Log) []string {
	row := []string{
		log.Service.ID,
		log.Request.Uri,
		log.Request.Url,
		fmt.Sprint(log.Request.Size),
		fmt.Sprint(log.Request.QueryString),
		log.Request.Headers.Accept,
		log.Request.Headers.Host,
		log.Request.Headers.UserAgent,
		log.UpStreamUri,
		fmt.Sprint(log.Response.Status),
		fmt.Sprint(log.Response.Size),
		log.Response.Headers.ContentLength,
		log.Response.Headers.Via,
		log.Response.Headers.Connection,
		log.Response.Headers.AccessControlAllowCredentials,
		log.Response.Headers.ContentType,
		log.Response.Headers.Server,
		log.Response.Headers.AccessControlAllowOrigin,
		log.AuthenticatedEntity.ConsumerID.Uuid,
		fmt.Sprint(log.Route.CreatedAt),
		log.Route.Hosts,
		log.Route.ID,
		fmt.Sprint(log.Route.Methods),
		fmt.Sprint(log.Route.Paths),
		fmt.Sprint(log.Route.PreserveHost),
		fmt.Sprint(log.Route.Protocols),
		fmt.Sprint(log.Route.RegexPriority),
		fmt.Sprint(log.Route.Service.ID),
		fmt.Sprint(log.Route.StripPath),
		fmt.Sprint(log.Route.UpdatedAt),
		fmt.Sprint(log.Service.ConnectTimeout),
		fmt.Sprint(log.Service.CreatedAt),
		fmt.Sprint(log.Service.Host),
		log.Service.Name,
		log.Service.Path,
		fmt.Sprint(log.Service.Port),
		log.Service.Protocol,
		fmt.Sprint(log.Service.ReadTimeout),
		fmt.Sprint(log.Service.Retries),
		fmt.Sprint(log.Service.UpdatedAt),
		fmt.Sprint(log.Service.WriteTimeout),
		fmt.Sprint(log.Latencies.Proxy),
		fmt.Sprint(log.Latencies.Kong),
		fmt.Sprint(log.Latencies.Request),
		log.ClientIp,
		fmt.Sprint(log.StartedAt),
	}

	return row
}

func getRowConvertedServiceTimeReport(logID string, requests int64, latencies *model.AverageLatencies) []string {
	row := []string{
		logID,
		fmt.Sprint(requests),
		fmt.Sprintf("%.2f", latencies.Proxy),
		fmt.Sprintf("%.2f", latencies.Kong),
		fmt.Sprintf("%.2f", latencies.Request),
		fmt.Sprintf("%.2f", latencies.Kong+latencies.Request+latencies.Proxy),
	}

	return row
}

func createCsvFile(outputName string, reportData [][]string) error {
	csvFile, err := os.Create(outputName)
	if err != nil {
		return err
	}

	csvwriter := csv.NewWriter(csvFile)

	for _, row := range reportData {
		_ = csvwriter.Write(row)
	}

	csvwriter.Flush()
	err = csvFile.Close()
	if err != nil {
		return err
	}

	return nil
}

func searchInArrayString(value string, values []string) bool {
	for i := range values {
		if value == values[i] {
			return true
		}
	}
	return false
}

func getIdLogs(kind string, logs []*model.Log) []string {
	var ids []string
	switch kind {
	case "consumer":
		for i := range logs {
			idConsumer := logs[i].AuthenticatedEntity.ConsumerID.Uuid
			if exists := searchInArrayString(idConsumer, ids); !exists {
				ids = append(ids, idConsumer)
			}
		}
	case "service":
		for i := range logs {
			idService := logs[i].Service.ID
			if exists := searchInArrayString(idService, ids); !exists {
				ids = append(ids, idService)
			}
		}
	}
	return ids
}

func createReportData(kind string, ids []string, logs []*model.Log) [][]string {
	var reportData [][]string
	listByID := map[string][]*model.Log{}

	//Criando um slice de logs para cada ID
	for i := range ids {
		listByID[ids[i]] = []*model.Log{}
	}

	switch kind {
	case "consumer":
		header := getConsumerReportHeader()
		reportData = append(reportData, header)
		//Separando dados por ID
		for j := range logs {
			logID := logs[j].AuthenticatedEntity.ConsumerID.Uuid
			listByID[logID] = append(listByID[logID], logs[j])
		}
		//Inserindo dados organizados por ID
		for i := range ids {
			for j := range listByID[ids[i]] {
				row := getRowConvertedConsumerReport(listByID[ids[i]][j])
				reportData = append(reportData, row)
			}
		}
	case "service":
		header := getServiceReportHeader()
		reportData = append(reportData, header)
		//Separando dados por ID
		for j := range logs {
			logID := logs[j].Service.ID
			listByID[logID] = append(listByID[logID], logs[j])
		}
		//Inserindo dados organizados por ID
		for i := range ids {
			for j := range listByID[ids[i]] {
				row := getRowConvertedServiceReport(listByID[ids[i]][j])
				reportData = append(reportData, row)
			}
		}
	}

	return reportData
}

func createServiceTimeReportData(ids []string, logs []*model.Log) [][]string {
	var reportData [][]string
	listByID := map[string][]*model.Log{}
	averageTimeByID := map[string]*model.AverageLatencies{}

	//Criando map dados por ID
	for j := range logs {
		logID := logs[j].Service.ID
		listByID[logID] = append(listByID[logID], logs[j])
	}
	for j := range ids {
		averageTimeByID[ids[j]] = &model.AverageLatencies{}
	}

	//Inserindo dados organizados por ID
	for i := range ids {
		for j := range listByID[ids[i]] {
			log := listByID[ids[i]][j]
			averageTimeByID[ids[i]].Kong += float64(log.Latencies.Kong)
			averageTimeByID[ids[i]].Proxy += float64(log.Latencies.Proxy)
			averageTimeByID[ids[i]].Request += float64(log.Latencies.Request)
		}
	}

	header := getServiceTimeReportHeader()
	reportData = append(reportData, header)

	//Somando dados de latência, kong e proxy por serviço
	for id, latencies := range averageTimeByID {
		logsPerService := int64(len(listByID[id]))
		latencies.Kong = float64(latencies.Kong) / float64(logsPerService)
		latencies.Proxy = float64(latencies.Proxy) / float64(logsPerService)
		latencies.Request = float64(latencies.Request) / float64(logsPerService)
		row := getRowConvertedServiceTimeReport(id, logsPerService, latencies)
		reportData = append(reportData, row)
	}

	return reportData
}

{{/* Expand the name of the chart */}}
{{- define "go-redis-sample.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/* Create chart name and version as used by the chart label */}}
{{- define "go-redis-sample.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/* Common labels */}}
{{- define "go-redis-sample.labels" -}}
helm.sh/chart: {{ include "go-redis-sample.chart" . }}
{{ include "go-redis-sample.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{/* Selector labels */}}
{{- define "go-redis-sample.selectorLabels" -}}
app.kubernetes.io/name: {{ include "go-redis-sample.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}} 
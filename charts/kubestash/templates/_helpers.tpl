{{/*
Expand the name of the chart.
*/}}
{{- define "kubestash.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "kubestash.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "kubestash.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "kubestash.labels" -}}
helm.sh/chart: {{ include "kubestash.chart" . }}
{{ include "kubestash.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "kubestash.selectorLabels" -}}
app.kubernetes.io/name: {{ include "kubestash.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "kubestash.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "kubestash.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Returns the appscode license
*/}}
{{- define "appscode.license" -}}
{{- default .Values.global.license .Values.license }}
{{- end }}

{{/*
Returns if the TaskQueue feature enabled or not
*/}}
{{- define "operator.enableTaskQueue" -}}
{{- default .Values.global.taskQueue.enabled .Values.taskQueue.enabled }}
{{- end }}

{{/*
{{- end }}

{{/*
Returns the maximum number of concurrent backupsessions
*/}}
{{- define "operator.maxConcurrentSessions" -}}
{{- default .Values.taskQueue.maxConcurrentSessions .Values.global.taskQueue.maxConcurrentSessions }}
{{- end }}

{{/*
Returns the appscode license secret name
*/}}
{{- define "appscode.licenseSecretName" -}}
{{- if .Values.licenseSecretName }}
{{- .Values.licenseSecretName -}}
{{- else if .Values.global.licenseSecretName }}
{{- .Values.global.licenseSecretName -}}
{{- else if (default .Values.global.license .Values.license) }}
{{- printf "%s-license" (include "kubestash.fullname" .) -}}
{{- end }}
{{- end }}

{{/*
Returns the registry used for operator docker image
*/}}
{{- define "operator.registry" -}}
{{- list (default .Values.registryFQDN .Values.global.registryFQDN) (default .Values.operator.registry .Values.global.registry) | compact | join "/" }}
{{- end }}

{{/*
Returns the registry used for kube-rbac-proxy docker image
*/}}
{{- define "rbacproxy.registry" -}}
{{- list (default .Values.registryFQDN .Values.global.registryFQDN) (default .Values.rbacproxy.registry .Values.global.registry) | compact | join "/" }}
{{- end }}

{{/*
Returns the registry used for catalog docker images
*/}}
{{- define "catalog.registry" -}}
{{- list (default .Values.registryFQDN .Values.global.registryFQDN) (default .Values.image.registry .Values.global.registry) | compact | join "/" }}
{{- end }}

{{/*
Returns the registry used for cleaner docker image
*/}}
{{- define "cleaner.registry" -}}
{{- list (default .Values.registryFQDN .Values.global.registryFQDN) (default .Values.cleaner.registry .Values.global.registry) | compact | join "/" }}
{{- end }}

{{/*
Returns whether the cleaner job YAML will be generated or not
*/}}
{{- define "cleaner.generate" -}}
{{- ternary "false" "true" (or .Values.global.skipCleaner .Values.cleaner.skip) -}}
{{- end }}

{{/*
Returns the appscode image pull secrets
*/}}
{{- define "appscode.imagePullSecrets" -}}
{{- with .Values.global.imagePullSecrets -}}
imagePullSecrets:
{{- toYaml . | nindent 2 }}
{{- else -}}
imagePullSecrets:
{{- toYaml $.Values.imagePullSecrets | nindent 2 }}
{{- end }}
{{- end }}

{{- define "image-pull-secrets" -}}
{{- $secrets:= list -}}
{{- with .Values.global.imagePullSecrets -}}
{{- range $x := . -}}
{{- $secrets = append $secrets $x.name -}}
{{- end -}}
{{- else -}}
{{- range $x := $.Values.imagePullSecrets -}}
{{- $secrets = append $secrets $x.name -}}
{{- end -}}
{{- end -}}
imagePullSecrets:
{{- toYaml $secrets | nindent 2 }}
{{- end -}}
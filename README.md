# Simple API Job Relay dengan keamanan JWT 

## Persyaratan yang dibutuhkan

1. Golang versi `1.18`
2. Module `gcc` untuk running sqlite3


## Cara menjalankan

```shell
simple-job-api$ go run main.go
2022/08/31 12:37:08 Starting & Bootstraping...

 ┌───────────────────────────────────────────────────┐ 
 │                   Fiber v2.37.0                   │ 
 │               http://127.0.0.1:3000               │ 
 │                                                   │ 
 │ Handlers ............. 7  Processes ........... 1 │ 
 │ Prefork ....... Disabled  PID ............. 15796 │ 
 └───────────────────────────────────────────────────┘ 
```

## Konfigurasi

Untuk mengatur port yang dipublish dan secret, Anda dapat melihat di `.env` file

Contoh `.env`

```
PORT=3000
SECRET=secret-changepls
```

## Dokumentasi API

### Register user

`POST http://localhost:3000/register`

Request body (json):
```json
{
	"name": "Wahyu",
	"username": "wahyuade",
	"password": "wahyuade"
}
```

Response sukses:
```json
{
	"status": true,
	"message": "Success"
}
```

### Login user

`POST http://localhost:3000/login`

Request body (json):
```json
{
	"password": "wahyuade",
	"username": "wahyuade"
}
```

Contoh response sukses:
```json
{
	"status": true,
	"message": "Success",
	"data": {
		"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjIxODEzNjMsImlkIjoiNGJlNDM0NzYtNzExYy00Y2VmLWFmZWUtOTc2MzMxNjAxNTc4IiwibmFtZSI6IldhaHl1IiwidXNlcm5hbWUiOiJ3YWh5dWFkZSJ9.cjR_Bdw3pHA-OuWG2UBY-HMU7aCRVncMxELYjR3Svm0"
	}
}
```

### Job List

`GET http://localhost:3000/jobs`

Pada metode ini saya hanya melakukan relay saja dan langsung diteruskan ke API Dans Multi Pro.

Gunakan `Authorization Bearer` dengan `token` dari response **Login user**
Implementasinya diletakkan pada header dengan ketentuan:

`Authorization: Bearer TOKEN_DARI_LOGIN_USER`

Parameters:

- `description`: `String` untuk search berdasarkan kata yang ada di description
- `location` : `String` untuk search berdasarkan location
- `full_time` : `Boolean` bernilai `true` / `false` (saya test dari API Dans Multi Pro tidak ada yang berubah ketika di set nilai `false`)
- `page` : `Number` digunakan untuk pagination


Contoh request:

`GET http://localhost:3000/jobs?description=golang`

Contoh response:

```json
{
	"status": true,
	"message": "Success",
	"data": [
		{
			"id": "0afc57b1-f535-4cf6-af94-4c1aa57bbcfa",
			"type": "Full Time",
			"url": "https://jobs.github.com/positions/0afc57b1-f535-4cf6-af94-4c1aa57bbcfa",
			"created_at": "Tue May 18 14:05:17 UTC 2021",
			"company_url": "https://bewerbungsportal.devk.de/index",
			"location": "Cologne",
			"title": "Cloud Platform Engineer (d/m/w)",
			"description": "<p>Cloud Platform Engineer (d/m/w)</p>\n<p>Seit mehr als 130 Jahren legen Versicherte in Deutschland ihre Risiken des Alltags vertrauensvoll in die Hände der DEVK Versicherungen. Heute ist die DEVK bundesweit präsent und betreut die persönlichen Anliegen von über vier Millionen Kunden in allen Versicherungssparten.\nIn der DEVK IT sind bereits heute Cloud Lösungen ein wesentlicher Baustein, der künftig weiter ausgebaut werden soll. Als Cloud Platform Engineer (d/m/w) sind Sie verantwortlich für die Operationalisierung der Cloud Strategie und die Befähigung der DEVK für den Einsatz von Cloud Services.</p>\n<p>Ihre Aufgaben</p>\n<ul>\n<li>Administration und Weiterentwicklung unserer Kubernetes Cluster</li>\n<li>Entwicklung und Betrieb der Infrastruktur-Automatisierung mit Terraform / Terragrunt / Terratest</li>\n<li>Optimierung der Deployment-Pipeline</li>\n<li>Implementierung und Administration AWS Services</li>\n<li>Betrieb, Entwicklung und Pflege von Docker Images</li>\n<li>Administration und Weiterentwicklung von Security-Komponenten in AWS und Kubernetes</li>\n</ul>\n<p>Ihr Profil</p>\n<ul>\n<li>Entwickler-Mindset (trunk-based development, continuous delivery) mit Erfahrungen im Bereich der Containerorchestrierung (bevorzugt Kubernetes)</li>\n<li>Erfahrungen mit Cloud Services (bevorzugt AWS)</li>\n<li>Breites technisches Knowhow im Bereich Cloud und Container</li>\n<li>Kenntnisse agiler Methoden</li>\n<li>Gute Kenntnisse der folgenden Sprachen, Tools oder Technologien:\no\tKubernetes, Helm\no\tHashicorp Stack (Vault, Terraform, Packer, Consul)\no\tGolang\no\tElasticSearch, Kibana</li>\n<li>Grundlagen in Netzwerk-Technologie</li>\n<li>Spaß an Teamarbeit (Kommunikation, Knowhow-Aufbau im Team etc.)</li>\n<li>Interesse an neuen Technologien und Offenheit für das Angebot der DEVK zur permanenten Weiterentwicklung</li>\n</ul>\n<p>Unsere Benefits</p>\n<ul>\n<li>Sichere Arbeitsplätze in einem stetig wachsenden Unternehmen</li>\n<li>Eine hervorragende arbeitgeberfinanzierte betriebliche Altersversorgung</li>\n<li>Flexible Arbeitszeitmodelle</li>\n<li>Möglichkeit für Homeoffice / Mobiles Arbeiten</li>\n<li>Stark vergünstigtes Job-Ticket</li>\n</ul>\n<p>Informieren Sie sich über unsere weiteren Benefits:\n<a href=\"https://www.devk.de/karriere/devkjobwelt/innendienst/mitarbeitervorteile-innendienst.jsp\">https://www.devk.de/karriere/devkjobwelt/innendienst/mitarbeitervorteile-innendienst.jsp</a></p>\n<p>Jetzt bewerben:\n<a href=\"https://bewerbungsportal.devk.de/jobs/15252-cloud-platform-engineer-d-m-w/job_application/new?cid=662c9407-f8cb-4024-a169-90c1c45dcfb4&amp;jid=7d3c50b7-4033-4607-a7e6-5a66a29062c0&amp;pid=3c460e75-1500-4628-862b-d5e5a9a02334&amp;it=4oCg1W3_iapbwK0VP9K01w\">https://bewerbungsportal.devk.de/jobs/15252-cloud-platform-engineer-d-m-w/job_application/new?cid=662c9407-f8cb-4024-a169-90c1c45dcfb4&amp;jid=7d3c50b7-4033-4607-a7e6-5a66a29062c0&amp;pid=3c460e75-1500-4628-862b-d5e5a9a02334&amp;it=4oCg1W3_iapbwK0VP9K01w</a></p>\n<p>DEVK-Stellenbörse:\n<a href=\"https://bewerbungsportal.devk.de/index\">https://bewerbungsportal.devk.de/index</a></p>\n<p>Wir freuen uns darauf, Sie persönlich kennenzulernen. Falls Sie noch Fragen haben, rufen Sie uns gerne an.</p>\n<p>Ihre Ansprechpartnerin:\nTanja Willing</p>\n<p>Tel.: +49 221 757-1154</p>\n",
			"how_to_apply": "<p><a href=\"https://bewerbungsportal.devk.de/jobs/15252-cloud-platform-engineer-d-m-w/job_application/new?cid=662c9407-f8cb-4024-a169-90c1c45dcfb4&amp;jid=7d3c50b7-4033-4607-a7e6-5a66a29062c0&amp;pid=3c460e75-1500-4628-862b-d5e5a9a02334&amp;it=4oCg1W3_iapbwK0VP9K01w\">https://bewerbungsportal.devk.de/jobs/15252-cloud-platform-engineer-d-m-w/job_application/new?cid=662c9407-f8cb-4024-a169-90c1c45dcfb4&amp;jid=7d3c50b7-4033-4607-a7e6-5a66a29062c0&amp;pid=3c460e75-1500-4628-862b-d5e5a9a02334&amp;it=4oCg1W3_iapbwK0VP9K01w</a></p>\n",
			"company_logo": "https://jobs.github.com/rails/active_storage/blobs/eyJfcmFpbHMiOnsibWVzc2FnZSI6IkJBaHBBaEtqIiwiZXhwIjpudWxsLCJwdXIiOiJibG9iX2lkIn19--24c86a54e3794712ee0bbdc4994243071b5bb978/DEVK.png"
		}
	]
}
```

### Job Detail

`GET http://localhost:3000/jobs/:id`

Untuk keterangan implementasi sama dengan **Job List**

Contoh request:

`GET http://localhost:3000/jobs/0afc57b1-f535-4cf6-af94-4c1aa57bbcfa`

Contoh response:

```json
{
	"status": true,
	"message": "Success",
	"data": {
		"id": "0afc57b1-f535-4cf6-af94-4c1aa57bbcfa",
		"type": "Full Time",
		"url": "https://jobs.github.com/positions/0afc57b1-f535-4cf6-af94-4c1aa57bbcfa",
		"created_at": "Tue May 18 14:05:17 UTC 2021",
		"company_url": "https://bewerbungsportal.devk.de/index",
		"location": "Cologne",
		"title": "Cloud Platform Engineer (d/m/w)",
		"description": "<p>Cloud Platform Engineer (d/m/w)</p>\n<p>Seit mehr als 130 Jahren legen Versicherte in Deutschland ihre Risiken des Alltags vertrauensvoll in die Hände der DEVK Versicherungen. Heute ist die DEVK bundesweit präsent und betreut die persönlichen Anliegen von über vier Millionen Kunden in allen Versicherungssparten.\nIn der DEVK IT sind bereits heute Cloud Lösungen ein wesentlicher Baustein, der künftig weiter ausgebaut werden soll. Als Cloud Platform Engineer (d/m/w) sind Sie verantwortlich für die Operationalisierung der Cloud Strategie und die Befähigung der DEVK für den Einsatz von Cloud Services.</p>\n<p>Ihre Aufgaben</p>\n<ul>\n<li>Administration und Weiterentwicklung unserer Kubernetes Cluster</li>\n<li>Entwicklung und Betrieb der Infrastruktur-Automatisierung mit Terraform / Terragrunt / Terratest</li>\n<li>Optimierung der Deployment-Pipeline</li>\n<li>Implementierung und Administration AWS Services</li>\n<li>Betrieb, Entwicklung und Pflege von Docker Images</li>\n<li>Administration und Weiterentwicklung von Security-Komponenten in AWS und Kubernetes</li>\n</ul>\n<p>Ihr Profil</p>\n<ul>\n<li>Entwickler-Mindset (trunk-based development, continuous delivery) mit Erfahrungen im Bereich der Containerorchestrierung (bevorzugt Kubernetes)</li>\n<li>Erfahrungen mit Cloud Services (bevorzugt AWS)</li>\n<li>Breites technisches Knowhow im Bereich Cloud und Container</li>\n<li>Kenntnisse agiler Methoden</li>\n<li>Gute Kenntnisse der folgenden Sprachen, Tools oder Technologien:\no\tKubernetes, Helm\no\tHashicorp Stack (Vault, Terraform, Packer, Consul)\no\tGolang\no\tElasticSearch, Kibana</li>\n<li>Grundlagen in Netzwerk-Technologie</li>\n<li>Spaß an Teamarbeit (Kommunikation, Knowhow-Aufbau im Team etc.)</li>\n<li>Interesse an neuen Technologien und Offenheit für das Angebot der DEVK zur permanenten Weiterentwicklung</li>\n</ul>\n<p>Unsere Benefits</p>\n<ul>\n<li>Sichere Arbeitsplätze in einem stetig wachsenden Unternehmen</li>\n<li>Eine hervorragende arbeitgeberfinanzierte betriebliche Altersversorgung</li>\n<li>Flexible Arbeitszeitmodelle</li>\n<li>Möglichkeit für Homeoffice / Mobiles Arbeiten</li>\n<li>Stark vergünstigtes Job-Ticket</li>\n</ul>\n<p>Informieren Sie sich über unsere weiteren Benefits:\n<a href=\"https://www.devk.de/karriere/devkjobwelt/innendienst/mitarbeitervorteile-innendienst.jsp\">https://www.devk.de/karriere/devkjobwelt/innendienst/mitarbeitervorteile-innendienst.jsp</a></p>\n<p>Jetzt bewerben:\n<a href=\"https://bewerbungsportal.devk.de/jobs/15252-cloud-platform-engineer-d-m-w/job_application/new?cid=662c9407-f8cb-4024-a169-90c1c45dcfb4&amp;jid=7d3c50b7-4033-4607-a7e6-5a66a29062c0&amp;pid=3c460e75-1500-4628-862b-d5e5a9a02334&amp;it=4oCg1W3_iapbwK0VP9K01w\">https://bewerbungsportal.devk.de/jobs/15252-cloud-platform-engineer-d-m-w/job_application/new?cid=662c9407-f8cb-4024-a169-90c1c45dcfb4&amp;jid=7d3c50b7-4033-4607-a7e6-5a66a29062c0&amp;pid=3c460e75-1500-4628-862b-d5e5a9a02334&amp;it=4oCg1W3_iapbwK0VP9K01w</a></p>\n<p>DEVK-Stellenbörse:\n<a href=\"https://bewerbungsportal.devk.de/index\">https://bewerbungsportal.devk.de/index</a></p>\n<p>Wir freuen uns darauf, Sie persönlich kennenzulernen. Falls Sie noch Fragen haben, rufen Sie uns gerne an.</p>\n<p>Ihre Ansprechpartnerin:\nTanja Willing</p>\n<p>Tel.: +49 221 757-1154</p>\n",
		"how_to_apply": "<p><a href=\"https://bewerbungsportal.devk.de/jobs/15252-cloud-platform-engineer-d-m-w/job_application/new?cid=662c9407-f8cb-4024-a169-90c1c45dcfb4&amp;jid=7d3c50b7-4033-4607-a7e6-5a66a29062c0&amp;pid=3c460e75-1500-4628-862b-d5e5a9a02334&amp;it=4oCg1W3_iapbwK0VP9K01w\">https://bewerbungsportal.devk.de/jobs/15252-cloud-platform-engineer-d-m-w/job_application/new?cid=662c9407-f8cb-4024-a169-90c1c45dcfb4&amp;jid=7d3c50b7-4033-4607-a7e6-5a66a29062c0&amp;pid=3c460e75-1500-4628-862b-d5e5a9a02334&amp;it=4oCg1W3_iapbwK0VP9K01w</a></p>\n",
		"company_logo": "https://jobs.github.com/rails/active_storage/blobs/eyJfcmFpbHMiOnsibWVzc2FnZSI6IkJBaHBBaEtqIiwiZXhwIjpudWxsLCJwdXIiOiJibG9iX2lkIn19--24c86a54e3794712ee0bbdc4994243071b5bb978/DEVK.png"
	}
}
```
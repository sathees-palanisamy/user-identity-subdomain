user-identity-subdomain

https://identity-f55axo5vlq-ez.a.run.app

gcloud builds submit --tag gcr.io/user-identity-subdomain-rest/identity


gcloud run deploy --image gcr.io/user-identity-subdomain-rest/identity --platform managed

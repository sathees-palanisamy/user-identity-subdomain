user-identity-subdomain

https://user-identity-subdomain-go.ew.r.appspot.com

gcloud builds submit --tag gcr.io/user-identity-subdomain-rest/identity


gcloud run deploy --image gcr.io/user-identity-subdomain-rest/identity --platform managed

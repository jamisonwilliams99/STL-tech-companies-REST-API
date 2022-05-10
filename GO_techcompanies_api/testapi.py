import requests

api_url = "http://127.0.0.1:8080/api/delete_company/1"

# type Company struct {
# 	Name           string `json:"name"`
# 	Industry       string `json:"industry"`
# 	Funding        string `json:"funding"`
# 	Employees      string `json:"employees"`
# 	EmployeeGrowth string `json:"employeegrowth"`
# 	Revenue        string `json:"revenue"`
# }

company = {"name": "Generic Company", "industry": "Software", "funding": "$8M", "employees": "10000", "employeegrowth": "55%", "revenue": "$12M"}

#response = requests.post(api_url, json=company)

response = requests.delete(api_url)

response.json()

response.status_code
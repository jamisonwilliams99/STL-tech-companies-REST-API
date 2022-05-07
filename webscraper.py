from bs4 import BeautifulSoup
import requests

def extract_name_from_url(url):
    idx = url.rindex("/")
    name = url[idx+1:]
    name = name.replace("_", " ")
    return name

url = "https://growjo.com/city/St_Louis"
api_url = "http://127.0.0.1:8080/api/create_company"

result = requests.get(url)

doc = BeautifulSoup(result.text, "html.parser")

tbody = doc.tbody
trs = tbody.contents

companies = []
for tr in trs:
    if len(tr.contents) == 1:
        continue
    
    company_name_td, industry_td, funding_td  = tr.contents[1], tr.contents[2], tr.contents[3]
    num_employees_td, employee_growth_td, revenue_td = tr.contents[4], tr.contents[5], tr.contents[6]

    company_name = company_name_td.a.text if company_name_td.a.text.find("...") == -1 else extract_name_from_url(company_name_td.a.get("href"))
    industry = industry_td.text
    funding = funding_td.text if funding_td.text != "" else "N/A"
    num_employees = num_employees_td.a.text
    employee_growth = employee_growth_td.text
    revenue = revenue_td.text

    company_data = {"name": company_name, "industry": industry, "funding": funding, "employees": num_employees, "employeegrowth": employee_growth, "revenue": revenue}
    companies.append(company_data)

for company in companies:
    response = requests.post(api_url, json=company)
    response.json()
    response.status_code






import requests
from bs4 import BeautifulSoup
import sys


url = f'https://ahmia.fi/search/?q={sys.argv[1]}'
reqs = requests.get(url)
soup = BeautifulSoup(reqs.text, 'html.parser')

urls = []
for link in soup.find_all('a'):
	print(link.get('href'))

import requests
from bs4 import BeautifulSoup
import ctypes
import numpy

def request(url):
    price_x = []
    price_y = []
    r = requests.get(url).text
    soup1 = BeautifulSoup(r, 'lxml')
    x = soup1.findAll("div", class_ = "js_td_width")

    soup2 = BeautifulSoup(r, 'lxml')
    y = soup2.findAll("div", class_ = "js_t_w")

    for y_iter in reversed(range(1, 4)):
        art_y = y[y_iter].text
        price_y.append(art_y.split(" ")[0].split(".")[0])

    price_y = list(map(numpy.float64, price_y))

    for i in reversed(range(0, 3)):
        art = x[i].text
        text = art.replace(" ", "")
        price_x.append(text)  
    
    price_x = list(map(numpy.float64, price_x))

    return price_x, price_y

def main():
    p_x, p_y = request("https://investfunds.ru/indexes/9021/")
    with open("parser.txt", "w") as f:
        for i in p_x:
            f.write(str(i)+" ")
        f.write("\n")
        for j in p_y:
            f.write(str(j)+" ")

if __name__ == "__main__":
    main()

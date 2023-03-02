import json

ori = open("./ori.txt", encoding="UTF-8")
st = ori.read()
d = {"info": st}
fin = json.dumps(d, ensure_ascii=False)
ori.close()
dst = open("./19.json", encoding="UTF-8",mode="x")
dst.write(fin)

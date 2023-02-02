# -- coding:utf-8 --
"""
1,海龟汤
2,谁是卧底
3,真心话大冒险
4,害你在心口难开
13,你演我猜（两人一组）
14,你演我猜（多人一组）
15,正话反说
16,动名词接龙
17,逛三园
18,你演我猜（两人一组）
19,你演我猜（多人一组）
20,正话反说
21,动名词接龙
22,逛三园
23,虎牌
24,吸管传输
25,踩气球
26,疯狂的字母
27,虎牌
28,吸管传输
29,踩气球
30,疯狂的字母
31,葫芦兄弟
32,心有千千结
33,击鼓传花
34,夹气球跑
35,葫芦兄弟
36,心有千千结
37,击鼓传花
38,夹气球跑
"""

while 1:
    line = input().split(",")
    if len(line) < 2:
        break
    line[1].encode("utf-8")
    file = open("{0}.html".format(line[0]), mode="w+")
    file.write(f"""<!DOCTYPE html>
<head>
    <meta charset="utf-8">
    <title>{line[0]}</title>
</head>""")
    file.close()

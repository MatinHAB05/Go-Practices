import sys

# sys.stdout = open("a.txt", "w")

persian_alphabet_map = {
    "ا": 0,
    "ب": 1,
    "پ": 2,
    "ت": 3,
    "ث": 4,
    "ج": 5,
    "چ": 6,
    "ح": 7,
    "خ": 8,
    "د": 9,
    "ذ": 10,
    "ر": 11,
    "ز": 12,
    "ژ": 13,
    "س": 14,
    "ش": 15,
    "ص": 16,
    "ض": 17,
    "ط": 18,
    "ظ": 19,
    "ع": 20,
    "غ": 21,
    "ف": 22,
    "ق": 23,
    "ک": 24,
    "گ": 25,
    "ل": 26,
    "م": 27,
    "ن": 28,
    "و": 29,
    "ه": 30,
    "ی": 31
}
persian_alphabet_reverse_map = {
    0: "ا",
    1: "ب",
    2: "پ",
    3: "ت",
    4: "ث",
    5: "ج",
    6: "چ",
    7: "ح",
    8: "خ",
    9: "د",
    10: "ذ",
    11: "ر",
    12: "ز",
    13: "ژ",
    14: "س",
    15: "ش",
    16: "ص",
    17: "ض",
    18: "ط",
    19: "ظ",
    20: "ع",
    21: "غ",
    22: "ف",
    23: "ق",
    24: "ک",
    25: "گ",
    26: "ل",
    27: "م",
    28: "ن",
    29: "و",
    30: "ه",
    31: "ی"
}
text = "«سی ذجو نط گقسوشی وز وسگوگ تپنپ ب جثسا جاجونز گوگی جن‌ذبگ. وز ذجو جن‌حبوینج سسگ وز جکاس طکگق ونق وسگوگ سی تبکا تسبگن، سگگ ضبجنق کو ظوه طقنگ.»"
reversed_text = text[-1:0:-1]
print("*"*50)
print("*"*50)
print("*"*50)
for s in range(0, 31):
    new_text = ""
    for x in text:
        if x in persian_alphabet_map:
            new_text += persian_alphabet_reverse_map[(
                s+persian_alphabet_map[x]) % 32]
        else:
            new_text += x

    print(new_text)
    print("*"*50)

print("*"*50)
print("*"*50)
print("*"*50)
for s in range(0, 31):
    new_text = ""
    for x in reversed_text:
        if x in persian_alphabet_map:
            new_text += persian_alphabet_reverse_map[(
                s+persian_alphabet_map[x]) % 32]
        else:
            new_text += x

    print(new_text)
    print("*"*50)

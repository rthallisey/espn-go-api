import json
import sys
import os

json_file = os.environ.get("JSON_DIR")
tmp_dir = os.path.dirname(json_file)

with open(json_file, "r") as f:
    json_string = f.read()
    obj = json.loads(json_string, encoding="utf-8")

for item in obj.keys():
#    with open(item.capitalize(), "w") as f:
    with open(tmp_dir + "/" + item.capitalize(), "w") as f:
        f.write("{\"" + item + "\": " + json.dumps(obj[item]) + "}")

    print item.capitalize()

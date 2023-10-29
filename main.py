import json
import math

data = []

centers = [
    {"x": 0, "y": 100},
    {"x": 5, "y": 70},
    {"x": 30, "y": 60},
    {"x": 90, "y": 50},
    {"x": 50, "y": 0},
    {"x": 5, "y": 5},
    {"x": 25, "y": 25},
    {"x": 75, "y": 25},
    {"x": 50, "y": 75},
    {"x": 100, "y": 100},
    {"x": 20, "y": 47},
    {"x": 50, "y": 50},
    {"x": 80, "y": 80},
    {"x": 100, "y": 100},
    {"x": 100, "y": 0},
]

for x in range(101):
    for y in range(101):
        maxZ = 0

        for center in centers:
            distanceToCenter = math.sqrt(
                math.pow(x - center["x"], 2) + math.pow(y - center["y"], 2)
            )
            z = math.floor((1 - distanceToCenter / 25) * 100)  # Adjusted divisor to make radii smaller
            if z > maxZ:
                maxZ = z

        data.append({"x": x, "y": y, "z": maxZ})

with open("result.json", "w") as file:
    json.dump(data, file)
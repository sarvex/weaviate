g.V().has("kind", "thing").has("uuid", "a079a15b-fa12-49d5-b4e8-5d82ec9449e3").as("class").V().has("uuid", "a079a15b-fa12-49d5-b4e8-5d82ec9449e3").optional(outE().has("refId").as("ref")).choose(select("ref"), select("class", "ref"), identity().project("class").by(select("class")))
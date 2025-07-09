package db

import "go.mongodb.org/mongo-driver/v2/bson"

var sampleData bson.A = bson.A{bson.M{"_id": "id_1", "content": "Insider - Project", "number": "+905555555555", "is_sent": false}, bson.M{"_id": "id_2", "content": "Insider - Project", "number": "+905555555555", "is_sent": false}, bson.M{"_id": "id_3", "content": "Insider - Project", "number": "+905555555555", "is_sent": false}}

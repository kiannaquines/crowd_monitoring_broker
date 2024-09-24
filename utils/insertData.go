package utils

import (
	"log"
)

func insertData() {
	
	log.Println("Inserting batch data...")

	tx, err := db.Begin()
	
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare(`
		INSERT INTO devices (device_id, device_addr, timestamp, is_randomized, device_power, ssid, frame_type, zone) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`)
	
	if err != nil {
		log.Fatal(err)
	}
	
	defer stmt.Close()

	for _, device := range deviceBuffer {
		_, err := stmt.Exec(device.UUID, device.DeviceAddr, device.Timestamp, device.IsRandomized, device.DevicePower, device.SSID, device.FrameType, device.Zone)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	deviceBuffer = deviceBuffer[:0]

	log.Println("Batch insert completed successfully")
}
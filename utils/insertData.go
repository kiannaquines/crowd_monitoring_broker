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
		INSERT INTO devices (device_addr, date_detected, is_randomized, device_power, frame_type, zone, processed, is_displayed) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	for _, device := range deviceBuffer {
		_, err := stmt.Exec(device.DeviceAddr, device.Timestamp, device.IsRandomized, device.DevicePower, device.FrameType, device.Zone, false, false)
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

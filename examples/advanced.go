package main

import (
	"fmt"
	"log"

	"github.com/jvherck/dicebear-go"
)

func main() {
	// Define custom options for the avatar
	options := &dicebear.Options{
		Flip:            true,                     // Flip the avatar horizontally
		Rotate:          90,                       // Rotate the avatar by 90 degrees
		Scale:           150,                      // Scale the avatar to 150%
		BackgroundColor: dicebear.Color("ffcc00"), // Set a custom background color
	}

	// Define custom parameters for the avatar
	customParams := map[string]string{
		"mouth":       "smile",   // Customize the mouth
		"eyes":        "happy",   // Customize the eyes
		"accessories": "glasses", // Add accessories
	}

	// Create a new avatar with the "avataaars" style, a custom seed, and options
	avatar, err := dicebear.NewAvatar(dicebear.Avataaars, "my-custom-seed", options, customParams)
	if err != nil {
		log.Fatalf("Failed to create avatar: %v", err)
	}

	// Print the avatar URL
	fmt.Println("Avatar URL:", avatar.URL())

	// Save the avatar as a PNG file
	outputPath := "avatar.png"
	_, err = avatar.Save(dicebear.PNG, outputPath, false)
	if err != nil {
		log.Fatalf("Failed to save avatar: %v", err)
	}

	fmt.Println("Avatar saved to:", outputPath)

	// Get the avatar schema (metadata)
	schema, err := avatar.GetSchema()
	if err != nil {
		log.Fatalf("Failed to get schema: %v", err)
	}

	fmt.Println("Avatar Schema:")
	for key, value := range schema {
		fmt.Printf("%s: %v\n", key, value)
	}
}

/*
Output:
Avatar URL: https://api.dicebear.com/9.x/avataaars/svg?seed=my-custom-seed&flip=true&rotate=90&scale=150&backgroundColor=ffcc00&mouth=smile&eyes=happy&accessories=glasses
Avatar saved to: avatar.png
Avatar Schema:
mouth: smile
eyes: happy
accessories: glasses
*/

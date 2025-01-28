package scenarios

import "artifactsmmo/pkg/api"

func GatherCopper(user api.ArtifactsUser) {
	user.ActionGathering("Azazel", true)

}

//result, err := user.ActionMove("Azazel", -1, 0)
//result, err := user.ActionFight("Azazel")

// for i := 0; i < 32; i++ {
// 	result, err := user.ActionGathering("Azazel")
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	cooldown, err := result.GetInt64("data", "cooldown", "remaining_seconds")
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	time.Sleep(time.Duration(cooldown) * time.Second)
// 	fmt.Printf("Slept for %d seconds, i = %d\n", cooldown, i)

// }

// fmt.Println("\nFinished execution!")

// result, err := user.ActionGathering("Azazel")

// if err != nil {
// 	log.Println(err)
// 	return
// } else {
// 	output := []byte(result.String())
// 	filename := fmt.Sprintf("log.json")
// 	err = os.WriteFile(filename, output, 0644)

// 	if err != nil {
// 		log.Panic(err)
// 	}
// }
// return

// for {
// 	result, err := user.ActionGathering("Azazel")
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	cooldown, err := result.GetInt64("data", "cooldown", "remaining_seconds")
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	time.Sleep(time.Duration(cooldown) * time.Second)
// 	fmt.Printf("Slept for %d seconds, time now: %s\n", cooldown, time.Now().Format(time.TimeOnly))

// 	inventory, _ := result.GetObjectArray("data", "character", "inventory")
// 	for _, item := range inventory {
// 		code, _ := item.GetString("code")
// 		quantity, _ := item.GetInt64("quantity")
// 		if (code == "copper_ore") && (quantity > 2) {
// 			result, err := user.ActionDeleteItem("Azazel", "copper_ore", int(quantity))
// 			if err != nil {
// 				log.Println(err)
// 				return
// 			}

// 			cooldown, err := result.GetInt64("data", "cooldown", "remaining_seconds")
// 			if err != nil {
// 				log.Println(err)
// 				return
// 			}

// 			time.Sleep(time.Duration(cooldown) * time.Second)
// 			fmt.Printf("Slept for %d seconds, time now: %s\n", cooldown, time.Now().Format(time.TimeOnly))
// 		}
// 	}

// }

// for i := 0; i <= 71; i++ {
// 	result, err := user.ActionGathering("Azazel")
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	cooldown, err := result.GetInt64("data", "cooldown", "remaining_seconds")
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	time.Sleep(time.Duration(cooldown) * time.Second)
// 	fmt.Printf("Completed %d iteration\n", i)

// }

// for {
// 	result, err := user.ActionFight("Azazel")
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	cooldown, err := result.GetInt64("data", "cooldown", "remaining_seconds")
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	time.Sleep(time.Duration(cooldown) * time.Second)

// 	result, err = user.ActionRest("Azazel")
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	cooldown, err = result.GetInt64("data", "cooldown", "remaining_seconds")
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	time.Sleep(time.Duration(cooldown) * time.Second)
// }

// for {
// 	// TODO add check that character is not here already!
// 	_, err := user.ActionMove("Azazel", 2, 0)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	user.WaitCooldown("Azazel")

// 	for i := 0; i <= 47; i++ {
// 		_, err = user.ActionGathering("Azazel")
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 		user.WaitCooldown("Azazel")
// 	}

// 	_, err = user.ActionMove("Azazel", 1, 5)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	user.WaitCooldown("Azazel")

// 	_, err = user.ActionCrafting("Azazel", "copper", 6)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	user.WaitCooldown("Azazel")

// 	_, err = user.ActionMove("Azazel", 2, 1)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	user.WaitCooldown("Azazel")

// 	_, err = user.ActionCrafting("Azazel", "copper_dagger", 1)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	user.WaitCooldown("Azazel")

// 	_, err = user.ActionRecycling("Azazel", "copper_dagger", 1)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	user.WaitCooldown("Azazel")

// 	_, err = user.ActionDeleteItem("Azazel", "copper", 2)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	user.WaitCooldown("Azazel")
// }

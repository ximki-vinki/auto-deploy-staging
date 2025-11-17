func ProcessServices(kuberNames []string, jsonConfig []byte) {
	// 1. Парсинг JSON в []Service
	var services []Service
	if err := json.Unmarshal(jsonConfig, &services); err != nil {
		log.Printf("Ошибка парсинга JSON: %v", err)
		return
	}

	// 2. Словарь для быстрого поиска
	serviceMap := make(map[string]Service)
	for _, s := range services {
		serviceMap[s.KuberName] = s
	}

	// 3. Проверка каждого kuberName
	for _, name := range kuberNames {
		s, exists := serviceMap[name]
		if !exists {
			log.Printf("Сервис %q: отсутствует", name)
			continue
		}
		if s.Disabled {
			log.Printf("Сервис %q: отключен", name)
		} else {
			log.Printf("Сервис %q: активен", name)
		}
	}
}
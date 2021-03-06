package clarifaiAPI

import "time"

type (
	Status struct {
		Code        uint   `json:"code"`
		Description string `json:"description"`
	}

	Concept struct {
		Id    string  `json:"id"`
		Name  string  `json:"name"`
		Value float64 `json:"value"`
		AppId string  `json:"app_id"`
	}

	Model struct {
		Name         string       `json:"name"`
		Id           string       `json:"id"`
		CreatedAt    time.Time    `json:"created_at"`
		AppID        string       `json:"app_id"`
		OutputInfo   OutputInfo   `json:"output_info"`
		ModelVersion ModelVersion `json:"model_version"`
		DisplayName  string       `json:"display_name"`
	}

	ModelVersion struct {
		Id        string    `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		Status    Status    `json:"status"`
	}

	OutputInfo struct {
		Message string `json:"message"`
		Type    string `json:"type"`
		TypeExt string `json:"type_ext"`
	}

	Output struct {
		Id          string      `json:"id"`
		Status      Status      `json:"status"`
		CreatedAt   time.Time   `json:"created_at"`
		Model       Model       `json:"model"`
		Input       Input       `json:"input"`
		ConceptData ConceptData `json:"data"`
	}

	OutputConfig struct {
		Language string `json:"language"`
	}

	OutputInf struct {
		OutputConfig OutputConfig `json:"output_config"`
	}

	Modl struct {
		OutputInf OutputInf `json:"output_info"`
	}

	ConceptData struct {
		Concepts []Concept `json:"concepts"`
	}

	Data struct {
		Image Image `json:"image"`
	}

	Image struct {
		URL string `json:"url"`
	}

	Input struct {
		Id   string `json:"id"`
		Data Data   `json:"data"`
	}

	Request struct {
		Inputs []Input `json:"inputs"`
		Modl   Modl    `json:"model"`
	}

	Response struct {
		Status  Status   `json:"status"`
		Outputs []Output `json:"outputs"`
	}
)

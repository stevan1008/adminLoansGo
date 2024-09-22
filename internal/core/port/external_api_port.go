package port

// Operaciones para integración de APIs externas (Canella, CRM, etc)
type ExternalAPIService interface {
	ValidateClientDocuments(clientID string) (bool, error)
	GetCreditScore(clientID string) (int, error) // Consulta el puntaje de crédito de un cliente con una agencia de crédito externa, de momento no hay implementación por el tiempo para estudiar un servicio gratuito para probar esta funcioonalida.
}

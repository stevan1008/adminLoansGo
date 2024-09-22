package service

import "github.com/stevan1008/adminLoansGo/internal/core/port"

type ExternalAPIServiceImpl struct {}

func NewExternalAPIService() port.ExternalAPIService {
    return &ExternalAPIServiceImpl{}
}

// Este apartado sirve para la validación de documentos de un cliente con una proveedor externo
func (s *ExternalAPIServiceImpl) ValidateClientDocuments(clientID string) (bool, error) {
    // Implementar la integración con la API de validación de documentos
    return true, nil
}

// Esto obtiene el puntaje de crédito del cliente desde una agencia externa
func (s *ExternalAPIServiceImpl) GetCreditScore(clientID string) (int, error) {
    // Esta implementación de momento queda vacia para usarse con un proveedor externo que se encargue de obtener el puntaje de un cliente para un credito
    return 700, nil
}
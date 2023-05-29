package handler

import (
	resultdto "dumbmerch/dto/result"
	transactiondto "dumbmerch/dto/transaction"
	"dumbmerch/models"
	"dumbmerch/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type HandlerTransactions struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *HandlerTransactions {
	return &HandlerTransactions{TransactionRepository}
}

func (h *HandlerTransactions) FindTransaction(c echo.Context) error {
	transaction, err := h.TransactionRepository.FindTransaction()

	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusOK,
			Message: "Waduh"})
	}
	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code: http.StatusOK,
		Data: transaction})

}

func (h *HandlerTransactions) FindTransactionId(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transaction, _ := h.TransactionRepository.FindTransactionId(id)

	if transaction.Id != id {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Data Gaada Bos"})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code: http.StatusOK,
		Data: transaction})
}

func (h *HandlerTransactions) DeleteTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transaction, _ := h.TransactionRepository.FindTransactionId(id)

	if transaction.Id != id {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "KOSONG OM"})
	}

	data, err := h.TransactionRepository.DeleteTransaction(id, transaction)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaction(data)})
}

func (h *HandlerTransactions) CreateTransaction(c echo.Context) error {
	request := new(transactiondto.CreateTransaction)
	// Id, _ := strconv.Atoi(c.Param("id"))
	// transaction, _ := h.TransactionRepository.FindTransactionId(Id)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	trips, err := h.TransactionRepository.GetTripId(request.IdTrip)

	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	transaction := models.Transaction{
		CounterQty: request.CounterQty,
		Total:      request.Total,
		Status:     "Waiting Approve",
		Attachment: request.Attachment,
		IdTrip:     request.IdTrip,
		Trip:       trips,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	data, err := h.TransactionRepository.CreateTransaction(transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaction(data)})
}

func (h *HandlerTransactions) UpdateTransaction(c echo.Context) error {
	request := new(transactiondto.UpdateTransaction)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	transaction, err := h.TransactionRepository.FindTransactionId(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	if request.CounterQty != 0 {
		transaction.CounterQty = request.CounterQty
	}
	if request.Total != 0 {
		transaction.Total = request.Total
	}
	if request.Status != "" {
		transaction.Status = request.Status
	}
	if request.Attachment != "" {
		transaction.Attachment = request.Attachment
	}
	if request.IdTrip != 0 {
		transaction.IdTrip = request.IdTrip
	}

	trips, err := h.TransactionRepository.GetTripId(request.IdTrip)

	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	transaction.Trip = trips

	transaction.UpdatedAt = time.Now()

	data, err := h.TransactionRepository.UpdateTransaction(id, transaction)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaction(data)})
}

func convertResponseTransaction(Transaction models.Transaction) transactiondto.TransactionResponse {
	return transactiondto.TransactionResponse{
		Id:         Transaction.Id,
		CounterQty: Transaction.CounterQty,
		Total:      Transaction.Total,
		Status:     Transaction.Status,
		Attachment: Transaction.Attachment,
		IdTrip:     Transaction.IdTrip,
		Trip:       Transaction.Trip,
	}
}

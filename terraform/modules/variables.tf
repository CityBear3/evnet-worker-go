variable "max_delivery_attempts" {
  description = "The maximum number of delivery attempts for dead letter policy"
  type        = number
  default     = 5
}

resource "google_pubsub_topic" "payment_method_created_topic" {
  name = "payment_method_created_topic"
}

resource "google_pubsub_topic" "payment_method_created_dead_later_topic" {
  name = "payment_method_created_dead_later_topic"
}

resource "google_pubsub_subscription" "payment_method_created_subscription" {
  name  = "payment_method_created_subscription"
  topic = google_pubsub_topic.payment_method_created_topic.name
  dead_letter_policy {
    dead_letter_topic = google_pubsub_topic.payment_method_created_dead_later_topic.id
    max_delivery_attempts = var.max_delivery_attempts
  }
  depends_on = [google_pubsub_topic.payment_method_created_topic, google_pubsub_topic.payment_method_created_dead_later_topic]
}

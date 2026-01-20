# PRD: Event Booking MVP

## 1. Introduction / Overview
Build an MVP for an event discovery and booking system. Users can search and view events (music, sports, films, festival, concert), book seats, complete checkout with a 15-minute seat hold, and receive tickets via email plus in-website notifications. Organizers can create, edit, and delete events. This MVP validates the core booking flow and operational readiness.

## 2. Goals
- Validate end-to-end booking for fixed-seat events (concerts/liveshows) with a 15-minute hold.
- Enable organizers to self-serve event creation and updates.
- Deliver tickets via email and on-site notifications reliably.
- Provide a clear audit trail for bookings, payments (card only), and notifications.

## 3. User Stories
- As a user, I can search for concerts/liveshow events and view details so I can decide what to attend.
- As a user, I can select a number of seats and optionally choose seat positions.
- As a user, I can reserve seats for 15 minutes while I complete checkout.
- As a user, I can confirm my contact information and payment method before paying.
- As a user, I receive my ticket by email and see a notification on the website.
- As an organizer, I can create a new event and publish it.
- As an organizer, I can edit event details or delete an event.
- As a user, I can request a refund, and the organizer can approve or reject it.

## 4. Functional Requirements
1. The system must allow users to search and view events.
2. The system must support fixed-seat events for the MVP, limited to event_type values: music, sports, films, festival, concert.
3. The system must show event details including date/time, venue, seating map (if available), price tiers, and availability.
4. The system must allow users to choose the quantity for each ticket type (price tier).
5. The system must allow users to select specific seat positions when a seating map is available for a fixed-seat tier.
6. The system must hold selected seats for 15 minutes during checkout.
7. The system must release held seats automatically after 15 minutes if payment is not completed.
8. The system must collect and confirm user contact information at checkout.
9. The system must support card payments only.
10. The system must confirm the payment method and proceed with payment.
11. The system must generate tickets upon successful payment based on booking items (quantity) and/or assigned seats.
12. The system must send the ticket to the user by email.
13. The system must send an in-website notification on booking success.
14. The system must allow organizers to create events (title, description, event_type, date/time, venue, seating map, pricing, total seats).
15. The system must allow organizers to edit events (while respecting existing bookings).
16. The system must allow organizers to delete events (with safeguards for existing bookings).
17. The system must allow users to request refunds for their bookings.
18. The system must allow organizers to approve or reject refund requests.
19. The system must record refund decision status and notify the user in-website and by email.
20. The system must log booking, payment, and notification outcomes for audit and support.
21. The system must model booking items with price tiers and quantities for general admission, and assigned seats for fixed-seat tiers.

## 5. Non-Goals (Out of Scope)
- General admission events and seatless inventory.
- Events outside music, sports, films, festival, and concert for this MVP.
- Additional payment methods (bank transfer, wallets, cash).
- SMS, push notifications, or messaging app notifications.
- Dynamic pricing, promo codes, or loyalty points.
- Resale/transfer of tickets.
- Multi-currency or multi-language support.

## 6. Design Considerations (Optional)
- Provide a clear, step-based checkout: seats -> info -> payment -> confirmation.
- Show a countdown timer for the 15-minute seat hold.
- Display seat availability with clear status (available, held, booked).
- Provide organizer forms with validation for dates, prices, and seat counts.

## 7. Technical Considerations (Optional)
- Integrate with existing Auth, Booking, Payment, Ticket, and Notification services.
- Ensure seat holds are atomic to prevent overselling (lock or reservation table).
- Model ticket types with an inventory_type per price tier to allow future mixed seating (fixed seats + general admission).
- Use idempotency for payment callbacks to avoid duplicate tickets.
- Trigger email and website notifications asynchronously after payment success.
- Maintain audit logs for bookings, payments, and refunds.

## 8. Success Metrics
- >= 80% of initiated checkouts either complete payment or release seats within 15 minutes without errors.
- <= 1% seat oversell incidents.
- >= 95% ticket email delivery success within 5 minutes of payment.
- Organizer can publish an event without support assistance.

## 9. Open Questions
- Refund policy details: time window, fees, and automatic vs manual steps.
- Which service owns seat maps and seat hold expiration logic?
- Should deleted events be hard-deleted or soft-deleted for audit?
- What exact ticket format is required (PDF, QR code, or text-only)?

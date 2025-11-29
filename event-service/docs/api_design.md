- [GET] `v1/events`
```
[
 {
   id: int
   name: string
   start_at: time
   end_at: time
   banner: string
   location
   status
 }
]
```
- [GET] `v1/events/{:id}`
```
{
 id: int
 name: string
 title: sttring
 start_at: time
 end_at: time
 banner: string
 location: string
 status: string
 ticket_types [
   {
     id: string
     position: int
     name: string
     description: string
     image_url: string
     status: string
     number_seats: int
     price: int
   }
 ]
}
```
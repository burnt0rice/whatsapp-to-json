# whatsapp-to-json
Converts exported Whatsapp chat txt file to json file.

## How to export the chat from whatsapp
1. Go to the chat
2. Click on the chats name
3. Scroll down and click on ``Export Chat``
4. Choose ``Without Media``
5. Save the .txt file to your computer

## Converting chat txt to json
Run ``main.go`` with the chat txt path. Replace ``_chat.txt`` with the file path of your chat file.
```shell
go run main.go -p "_chat.txt"
```

After running the code, there will be a new file called ``chat.json`` in the project directory.

## Example of the generated json file
```json
{
  "data": [
    {
      "date": "05.12.16",
      "time": "10:31:33",
      "sender": "Marc",
      "message": "Hello my friend"
    },
    {
      "date": "05.12.16",
      "time": "10:32:31",
      "sender": "Nico",
      "message": "Hello"
    },
    {
      "date": "05.12.16",
      "time": "10:35:21",
      "sender": "Marc",
      "message": "This is funny"
    }
  ]
}
```

### install pip
```
pip install telethon

https://api.telegram.org/botAPITOKENNUMBER:APITOKENKEYHERE/sendmessage?chat_id=-100GROUPNUMBER&text=test

https://api.telegram.org/botXXX:YYYY/getUpdates

https://api.telegram.org/bot<YourBOTToken>/getUpdates

curl https://api.telegram.org/bot<TOKEN>/getUpdates | jq
 
https://api.telegram.org/bot123456789:jbd78sadvbdy63d37gda37bd8/getUpdates

curl -X POST "https://api.telegram.org/botXXX:YYYY/sendMessage" -d "chat_id=-zzzzzzzzzz&text=my sample text"

```

### telegram 1
```
from telethon import InteractiveTelegramClient
from telethon.utils.tl_utils import get_display_name

client = InteractiveTelegramClient('session_id', 'YOUR_PHONE_NUMBER', api_id=1234YOURAPI_ID, api_hash='YOUR_API_HASH')

dialog_count = 10
dialogs, entities = client.get_dialogs(dialog_count)
for i, entity in enumerate(entities):
                    i += 1  # 1-based index
                    print('{}. {}. id: {}'.format(i, get_display_name(entity), entity.id))
```
### telegram 2
```
import os
import sys
from telethon import TelegramClient
from telethon.utils import get_display_name

import nest_asyncio
nest_asyncio.apply()

session_name = "<session_name>"
api_id = <api_id>
api_hash = "<api_hash>"
dialog_count = 10 # you may change this

if f"{session_name}.session" in os.listdir():
    os.remove(f"{session_name}.session")

client = TelegramClient(session_name, api_id, api_hash)

async def main():
    dialogs = await client.get_dialogs(dialog_count)
    for dialog in dialogs:
        print(get_display_name(dialog.entity), dialog.entity.id)

async with client:
    client.loop.run_until_complete(main())
```

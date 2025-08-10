Facebook Groups Auto-Poster (GraphQL)

Post to 100+ Facebook Groups in one click using your existing account — no weird logins, no APIs, just your browser session.


---

What It Does

Finds all your joined groups automatically

Uploads multiple photos per post

Reads your content (text + images) from a folder

Posts to groups with random delays for safety

Uses your own Facebook session via “Copy as cURL”



---

Quick Start

1. Install

git clone <repository-url>
cd POST_WITH_GRAPHQL
go mod tidy

2. Prepare Your Content

Folder format:

CONTENT_ROOT/
├── item1/
│   ├── details.txt    # description: Your text...Feature one...Feature two
│   ├── image1.jpg
│   └── image2.jpg
├── item2/
│   ├── details.txt
│   └── image1.jpg

3. Get Your cURL Strings

1. Open Facebook in your browser (logged in).


2. Open DevTools → Network.


3. Perform these actions:

View your groups (for fetchGroupsCurl)

Start a post with images (for uploadImageCurl)

Post to a group (for createPostCurl)



4. Right-click each request → Copy → Copy as cURL.


5. Paste them into main.go.



4. Run

make run
# or
go run main.go


---

Tips

Use fresh cURL copies (tokens expire quickly)

Keep delays between posts for account safety

Make sure all cURL requests are from the same browser session



---

📞 Support

Email: aronkipkorir254@gmail.com

WhatsApp: 0701416017.
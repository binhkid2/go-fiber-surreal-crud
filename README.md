# go-fiber-surreal-crud
Make sure you’ve installed SurrealDB — it should only take a second!
Start the database
macOS or Linux
user@localhost % surreal start memory -A --auth --user root --pass root
Windows
PS C:\> surreal.exe start memory -A --auth --user root --pass root
Once you have your database running, head over to #https://surrealist.app/. Before you can start using Surrealist you will have to input the credentials for your session,


<div>
 <p>
    EndpointURL: 'http://127.0.0.1:8000/'
 </p>
 <p>Namespace: 'test',</p>
<p>Database: 'test',</p>
<p>Username: 'root',</p>
<p>Password: 'root',</p>
</div>
# add some example query like </br>
<code>
CREATE biolinks SET
	title = 'title 1',
	link = "link1",
 isPublic = false
;
CREATE biolinks SET
	title = 'title 2',
	link = "link2",
 isPublic = true
;

  
</code>
### You are good to go
<code>
  go run .
</code>


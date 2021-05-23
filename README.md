# uconfp
Universal configuration file parser utility

Usage: uconfp [OPTIONS] COMMAND KEY [VALUE] FILENAME
Universal configuration file parser

Options:
  -t, --type		   Configuration file format (Available formats: INI)
  -s, --section section   Section name inside INI file to parse configuration (eg: [global])
  -d, --delimiter         Config key/value delimiter. (eg: key=value) (default: "=")
   q:q
Commands:
  get		Get configuration key value (eg: uconfp get -s mysqldump max_allowed_packet /etc/crontab)
  set		Set configuration key value (eg: uconfp set SHELL /bin/bash /etc/crontab)
  del		Delete configuration key (eg: uconfp del SHELL /bin/bash
  
Run uconfp COMMAND --help for more information on a command.


  
	

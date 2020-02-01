Job Applicant Assistant
========================
This is a web server using Gin-Gonic framework, and serve as a chrome extension.

Chrome extension link: <a href="https://tinyurl.com/y4j54rpl">求職助手</a>
------------------------

The chrome extension also open soursce at there <a href="https://github.com/CaiYueTing/job-apply-assistant">job apply assistant</a>

Deploy with <a herf="https://github.com/apex/up">apex/up</a>
------------------------
```bash
go mod download
mv up.json.env up.json
modify your up setting
up -v
```

Demo Service
---------------------------
The service will be matching the 104jobbank website, and the extension application will open.

<h3>Match Url<br>

![alt text](https://github.com/CaiYueTing/Job_Applicant_Assistant/blob/master/demo/match_url.gif)

<h3>Illegal Record<br>

![alt text](https://github.com/CaiYueTing/Job_Applicant_Assistant/blob/master/demo/illegalrecord.gif)

<h3>Qollie Link<br>

![alt text](https://github.com/CaiYueTing/Job_Applicant_Assistant/blob/master/demo/qollie.gif)

<h3>Welfare Abstract<br>

![alt text](https://github.com/CaiYueTing/Job_Applicant_Assistant/blob/master/demo/welfare.gif)

<h3>Salary<br>

![alt text](https://github.com/CaiYueTing/Job_Applicant_Assistant/blob/master/demo/salary.gif)
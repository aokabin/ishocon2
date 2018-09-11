.PHONY : rotate
rotate:
	sudo mv /var/log/nginx/access.log /var/log/nginx/access.log.`date +%s`
	sudo touch /var/log/nginx/access.log
	sudo systemctl restart nginx.service # 利用状況に応じて


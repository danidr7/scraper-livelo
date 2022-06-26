## Consultar pontuação
curl 'https://apis.pontoslivelo.com.br/partners-campaign/v1/campaigns/active?partnersCodes=SPR,AMZ,BOK,CEA,DPS,MOV,PTF,RNN,RSC' \
  -H 'authority: apis.pontoslivelo.com.br' \
  -H 'sec-ch-ua: "Google Chrome";v="95", "Chromium";v="95", ";Not A Brand";v="99"' \
  -H 'accept: application/json, text/javascript, */*; q=0.01' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36' \
  -H 'sec-ch-ua-platform: "macOS"' \
  -H 'origin: https://www.livelo.com.br' \
  -H 'sec-fetch-site: cross-site' \
  -H 'sec-fetch-mode: cors' \
  -H 'sec-fetch-dest: empty' \
  -H 'referer: https://www.livelo.com.br/' \
  -H 'accept-language: pt-BR,pt;q=0.9,en-US;q=0.8,en;q=0.7' \
  --compressed

## Consultar todos os partners
curl 'https://www.livelo.com.br/ccstore/v1/files/thirdparty/config_partners_compre_e_pontue.json' \
  -H 'sec-ch-ua: "Google Chrome";v="95", "Chromium";v="95", ";Not A Brand";v="99"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'Authorization: Bearer null' \
  -H 'Accept: application/json, text/javascript, */*; q=0.01' \
  -H 'Referer: https://www.livelo.com.br/ganhe-pontos-compre-e-pontue' \
  -H 'X-Requested-With: XMLHttpRequest' \
  -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36' \
  -H 'sec-ch-ua-platform: "macOS"' \
  --compressed
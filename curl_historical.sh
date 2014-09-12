#!/bin/bash



curl "http://www.nasdaq.com/symbol/$symbol/historical" -H 'Pragma: no-cache' -H 'Origin: http://www.nasdaq.com' -H 'Accept-Encoding: gzip,deflate' -H 'Accept-Language: en-US,en;q=0.8' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/37.0.2062.94 Safari/537.36' -H 'Content-Type: application/json' -H 'Accept: */*' -H 'Cache-Control: no-cache' -H "X-Requested-With: XMLHttpRequest' -H 'Cookie: c_enabled$=true; netseer_cm=done; userSymbolList=+; userCookiePref=true; fsr.s=; clientPrefs=||||lightg; s_pers=; s_sess=; NSC_W.TJUFEFGFOEFS.OBTEBR.80=ffffffffc3a08e3245525d5f4f58455e445a4a423660" -H 'Connection: keep-alive' -H "Referer: http://www.nasdaq.com/symbol/$symbol/historical" -H 'DNT: 1' --data-binary "10y|false|$symbol" --compressed

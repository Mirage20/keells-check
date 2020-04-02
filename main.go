package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gen2brain/beeep"
	"github.com/mirage20/keells-check/signals"
	"log"
	"net/http"
	"strings"
	"time"
)

var x = `
<!DOCTYPE html>
<html>
<head><title>
	Welcome to Keells Super-The First Online Supermarket in Sri Lanka
    
</title><meta charset="utf-8" /><meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, shrink-to-fit=no" /><meta http-equiv="Content-Type" content="text/html; charset=utf-8" /><meta name="msvalidate.01" content="132634E2E230F7A6660FC2854FC4A1EF" /><meta http-equiv="imagetoolbar" content="false" /><meta name="author" content="Keells Super" /><meta name="image" content="https://www.keellssuper.com/Themes/Images/ks_fb_share_new.jpg" /><meta property="og:title" content="The First Online Supermarket in Sri Lanka - www.keellssuper.com" /><meta property="og:type" content="website" /><meta property="og:description" content="The First Online Supermarket in Sri Lanka - www.keellssuper.com" /><meta property="og:url" content="https://www.keellssuper.com" /><meta property="og:site_name" content="www.keellssuper.com" /><meta property="og:image" content="https://www.keellssuper.com/Themes/Images/ks_fb_share_new.jpg" />
<link rel="icon" href="https://www.keellssuper.com/Themes/images/keellsfavicon.ico?v=2" />

<link href="https://fonts.googleapis.com/css?family=Source+Sans+Pro:300,400,600,700,900" rel="stylesheet" />

<link rel="stylesheet" href="/Themes/bootstrap.css" />

<link rel="stylesheet" href="/Themes/bootstrap_custom.css" />


<link rel="stylesheet" href="/Themes/vendors/menu/css/menu.css">

<link rel="stylesheet" href="/Themes/vendors/carousel/css/owl.carousel.css">
<link rel="stylesheet" href="/Themes/vendors/carousel/css/owl.theme.default.css">

<link rel="stylesheet" href="/Themes/vendors/scrollbars/css/scrollbar.css">

<link rel="stylesheet" href="/Themes/vendors/popup/css/popup.css">

<link rel="stylesheet" href="/Themes/vendors/scroll-header/css/scroll-header.css">

<link rel="stylesheet" href="/Themes/vendors/scroll-page/css/scroll-page.css">

<link rel="stylesheet" href="/Themes/vendors/slick/css/slick.css">
<link rel="stylesheet" href="/Themes/vendors/slick/css/slick-theme.css">
<link rel="stylesheet" href="/Themes/vendors/fancy/css/jquery.fancybox.css">

<link rel="stylesheet" href="/Themes/vendors/form/radio/css/md-radio.css">
<link href="/Styles/ui-lightness/jquery-ui-1.8.14.custom.css" rel="stylesheet" type="text/css" />
<script src="/Scripts/jquery-3.3.1.min.js"></script>
<script src="/Scripts/popper.min.js"></script>
<script src="/Scripts/bootstrap.min.js"></script>
<script src="/Scripts/jquery-ui-1.12.1.min.js"></script>
<script src="/Themes/vendors/popup/js/popup-animation.js"></script>
<script src="/Themes/vendors/popup/js/velocity.min.js"></script>
<script src="/Themes/vendors/popup/js/velocity.ui.min.js"></script>
<script src="/Themes/vendors/popup/js/quantity-input.js"></script>
<script src="/Scripts/jquery.validate.min.js"></script>
<script src="/Scripts/jquery.session.js"></script>
<script src="/Scripts/commonscripts.js"></script>
</head>
<body>
<form method="post" action="./login" id="form1">
<input type="hidden" name="__VIEWSTATE" id="__VIEWSTATE" value="EsE+GXFppSex0TKpmW0zu/fw2/JTELuH+APN7NmA52ScR5FVh3Swt+zxjpWiLbHLCDa8eN7tjg6VVt/Ait5oxMJPw1l6IAG3JJ6SY7GSnzXeDu5av2Pe1GcY0me81KxylQ8AJQ0l6v5zEvzlODzExV7ULwT7eqR5fQMIVhEEJcCDRrC6Irp9kVOPD+qtq87SJwzhoSnbZwv6g/KaxJsbqt2GqSlgqKL/8JbLACtYNNXtGpsezWB2XUSYFzQ2qdZg+otVDpuK7vmZXbry+ahRyKeSs+YeK5XjAieET6ZT7IlX8dQLq0lm1k01mZpIUfiwwWTHkESrUcNz9D/n7+CyFw7xyag38CsUbZT5cfrl9sX+i95poNJpZ7dK45+k1mbtth46f9zPHOFQagmAhqppcM1OZ8WB/um84HuW7eF56gGOkqTGH9lvYvCcqmruWnZ09slp47OtAjsp4Kpkjs4yeMrlO4V52Ah+NSmeWiVL8cw6NX4oLFJhKS+WNFDQ35eYqOy2Xzyciq3SBcdcUEiTx2Fq+XqR9/2yIUHoC2iW6rbYD6WoafGloQkRJyaweryoYsAqDRvMxgJ705EEn2QSXuGbMPCP+pj+HSL/m6DrcscMmDW2awZ+UUSOddhGz8zn6zoGhPtiZzO8OeuQhVb/ioesep0MjT35ok5xl5jbVKTDcoQEJiDR7zX3w8miYScLdP7UK1l33BVDhYB+wvRhZa89ORC2BmTyUke7qgt14dpT8SaaGzYosdy2XBlITtkAWfEUUFZE3iQzQSUySgmE/Dl1I7BmBhUSlr3vjAwNIU9+3Gz5h7wgxS7dNHWGdGI/JBHFHINUw2Py5fF90h9Jg1c7dANB9JJuwuc59cZy7oklj1yYDlj95HvZ9n1qVsOnhMhs3ItzOVyBl3MXRXiQsyQjRzdAARG7to4tWtKP1J5IgSTSBvbPdMO2YtixZUmTpX+lnK8luJcDKBVcU8elPw==" />
<input type="hidden" name="__VIEWSTATEGENERATOR" id="__VIEWSTATEGENERATOR" value="CA0B0334" />
<input type="hidden" name="__EVENTVALIDATION" id="__EVENTVALIDATION" value="vooEBDLBhzRNbhXtCXdbwYj975dWlIgxaoUJWTmzRRFzoFi5v6lKIzs/O7hOGGC/VbLlX+sbK8Af3lo2QK0jD5RfXgGpXy2bv3iUmAESkdD6gthN+B0UQzpuSUpNKw5Y1PfDkM8gRhYUva3Gq3fQwF3L9LJ5RWlDbvRoSZro9jIRqW0ucD1WRbGUAiX6DlKl+hkNWsrbVzptE6JwntDzUGabqFGtgkHWyv0WpernqDWYu8ct0OxnOaAlUg3SyW4dmV4wQkFbfUfo1g6hMT+Zn2RwzVacdcCE6VkYMCwFGNGBUD+2DdCHcUlsf2C8d1F8/RxfqOQnCyNeqAeWVp1Ay72+qPwWB4mW63t6vgkiOach2rgP+qs40oHub+hEnjHpbfpdthBPRZZmM6+sKwQ7TA2bcRnf8ddUQA/pZLO4DaGhXAsbL/uU0hSXO9bCDK+m7/0JXncfxCUNLjmZ7eK5yUCbGj6brFz40TgTffTQEuEvOsvJW1tGhFsO5KMgCxpGPbQllXvMW1UsMTJ7Zc73YOg0mIfo6u4Fl/QUM0NrjOD8TXUIBolU3dBeLDWI7K6dJLEF5NkWSs1gMygyKGeRwdWv1rpiS233BnLHKsuA608PFjj9ef4ErIpgSxgVeXPk" />
<div>
<div id="content" class="templogin container">
<div class="signInPage">
<div class="temploginlogo">
<img src="Themes/images/Logo-Keells.png" />
</div>
<div class="row">
<div class="page-title col-12">
<h1>Welcome to Keells Online</h1>
</div>
</div>

<div class="deliveryType-page">
<div class="styleBox pt-lg-5 pb-lg-4">
<div class="row justify-content-md-center">

<div class="col-lg-5 col-md-6 mb-4 changeCity-left" style="text-align: center;">
<div style="visibility: hidden; display: none;">
<span id="BodyContent_RblSelectOptions" class="md-radio rounded right outline-gray btn-block"><input id="BodyContent_RblSelectOptions_0" type="radio" name="ctl00$BodyContent$RblSelectOptions" value="Home" checked="checked" /><label for="BodyContent_RblSelectOptions_0">Home Delivery (Within Delivery Grid)</label></span>
</div>
<input type="hidden" name="ctl00$BodyContent$hfredirecturl" id="BodyContent_hfredirecturl" />
<h6>
<b><span style="color:red;">If your delivery area is not listed in our delivery list, it could be because</span><br /></b>
<span>
• The quota allocated for your area has been fulfilled for today or</span><br /><span>
•We are not delivering to your area yet.</span><br /><span>
We kindly request you not to login to the site if your delivery areas is not visible and allow another customer to do so. We regret any inconvenience caused and thank your for your patience.</span>
</h6>
<input type="image" name="ctl00$BodyContent$ImgOulets" id="BodyContent_ImgOulets" src="Themes/images/WebButtons-01.png" onclick="javascript:WebForm_DoPostBackWithOptions(new WebForm_PostBackOptions(&quot;ctl00$BodyContent$ImgOulets&quot;, &quot;&quot;, true, &quot;&quot;, &quot;&quot;, false, false))" style="width:59%" />
</div>

<div class="col-lg-7 col-md-6 pl-lg-5 changeCity-right">
 <div id="BodyContent_HomeDelivery" class="row">
<div class="col-md-12 text-center">
<p>
Please select your delivery area.<br>
Prices of some perishable items will depend on the delivery area.
</p>
</div>
<div class="col-md-12">
<div id="BodyContent_cityBox" class="row">
<div class="col-md-6">
<div class="form-group">
<label>Delivery City</label>
<select name="ctl00$BodyContent$ddlDeliveryCity" id="BodyContent_ddlDeliveryCity" class="custom-select">
<option value="326">Chilaw</option>
<option value="263">Galle</option>
<option value="228">Kurunegala</option>
<option value="329">Marawila</option>
<option value="282">Matale</option>
<option value="264">Matara</option>
<option value="331">Puttalam</option>
<option value="273">Rathnapura</option>
<option value="303">Wennappuwa</option>
<option selected="selected" value="-1">---select---</option>
</select>
</div>
</div>
<div class="col-md-6">
<div class="form-group">
<label>Suburb</label>
<select name="ctl00$BodyContent$ddlSuburb" id="BodyContent_ddlSuburb" class="custom-select">
<option selected="selected" value="-1">---select---</option>
</select>
</div>
</div>
<div class="message" style="display: none">
<span id="BodyContent_lblErrorMsg"></span>
</div>
</div>
<div class="form-group" style="display: none;">
<label>Search Delivery Area</label>
<div class="form-group has-search">
<div id="BodyContent_pnlChangeCity">
<span class="icon icon-search2 form-control-feedback">
<a id="BodyContent_lbtnAutoSubSearch" href="javascript:WebForm_DoPostBackWithOptions(new WebForm_PostBackOptions(&quot;ctl00$BodyContent$lbtnAutoSubSearch&quot;, &quot;&quot;, true, &quot;&quot;, &quot;&quot;, false, true))"></a>
</span>
<input name="ctl00$BodyContent$txtAutoSuburb" type="text" id="BodyContent_txtAutoSuburb" class="form-control" placeholder="Search Suburb" autocomplete="off" />
<span id="BodyContent_lblAutoSuburbSearch"><font color="Red"></font></span>
<div style="visibility: hidden; display: none">
<input type="submit" name="ctl00$BodyContent$BtnHiddenSelectedSuburb" value="Button" onclick="javascript:WebForm_DoPostBackWithOptions(new WebForm_PostBackOptions(&quot;ctl00$BodyContent$BtnHiddenSelectedSuburb&quot;, &quot;&quot;, true, &quot;&quot;, &quot;&quot;, false, false))" id="BodyContent_BtnHiddenSelectedSuburb" />
<input type="hidden" name="ctl00$BodyContent$hfSelectedSuburb" id="BodyContent_hfSelectedSuburb" />
</div>

</div>
</div>
<div id="BodyContent_Searchsuburblist" class="suburblist">
<ul class="ul-suburblist">
</ul>
</div>
<div class="pop">
</div>
</div>
</div>
</div>
<div style="display: none;">
<div class="form-group text-center">
<input type="submit" name="ctl00$BodyContent$BtnContinueShopping" value="Continue Shopping" id="BodyContent_BtnContinueShopping" class="btn btn-outline-primary" />
</div>
</div>
</div>
</div>
</div>
<div class="row">
<div class="styleBox sign-in-block">
<div class="page-title col-12">
<p>If you shopped with us before, please proceed to login by entering your user name and password.</p>
<p>If you dont have an account, register now.</p>
<div style="                                    display: none;
                            ">
<span>Currently logged in users </span>67 out of 70
</div>
</div>
<p>Sign in</p>
<div class="row">
<div class="form-group col-sm-12" style="text-align:center;">
</div>
</div>
<div class="row">
<div class="form-group col-sm-6">
<label for="exampleInputEmail1">Email address</label>
<input name="ctl00$BodyContent$UserName" type="text" maxlength="50" id="BodyContent_UserName" class="form-control" placeholder="Enter email" />
</div>
<div class="form-group col-sm-6">
<label for="exampleInputPassword1">Password</label>
<input name="ctl00$BodyContent$LoginPassword" type="password" maxlength="31" id="BodyContent_LoginPassword" class="form-control" autocomplete="off" placeholder="Password" />
</div>
<div class="col-12">
<div class="forgot-pass">
<a href='https://int.keellssuper.net/forgot-password'>Forgot password</a>
</div>
</div>
<div class="col-sm-12 text-sm-center mob-width">
<input type="submit" name="ctl00$BodyContent$BtnRegister" value="New User Registration" onclick="javascript:WebForm_DoPostBackWithOptions(new WebForm_PostBackOptions(&quot;ctl00$BodyContent$BtnRegister&quot;, &quot;&quot;, true, &quot;&quot;, &quot;&quot;, false, false))" id="BodyContent_BtnRegister" class="btn btn-primary" />
<input type="submit" name="ctl00$BodyContent$BtnLogin" value="Sign in" onclick="ga(&#39;send&#39;, &#39;event&#39;, &#39;Click&#39;, &#39;Login button&#39;);WebForm_DoPostBackWithOptions(new WebForm_PostBackOptions(&quot;ctl00$BodyContent$BtnLogin&quot;, &quot;&quot;, true, &quot;LoginUserValidationGroup&quot;, &quot;&quot;, false, false))" id="BodyContent_BtnLogin" class="btn btn-primary" />
</div>
</div>
</div>
</div>
</div>
<div id="LoginDialog" class="modal fade bd-example-modal-md styleBox-popup" data-easein="bounce" tabindex="-1" role="dialog" aria-labelledby="myLargeModalLabel" aria-hidden="true">
<div class="modal-dialog modal-md pt-4">
<div class="modal-content">
<div class="row">
<div class="col-md-12">
<div class="iconTitleBlock">
<div class="icon-box">
<img src="/Themes/images/icon-danger-green.svg">
</div>
Important!
</div>
</div>
<div class="col-md-12 text-center">
<span id="BodyContent_lblShowMessage"></span>
</div>
<div class="col-md-12">
<div class="modal-footer justify-content-center">
<input type="submit" name="ctl00$BodyContent$btnSubmit" value="OK" onclick="javascript:WebForm_DoPostBackWithOptions(new WebForm_PostBackOptions(&quot;ctl00$BodyContent$btnSubmit&quot;, &quot;&quot;, true, &quot;&quot;, &quot;&quot;, false, false))" id="BodyContent_btnSubmit" class="btn btn-outline-primary" data-dismiss="modal" />
</div>
</div>
</div>
</div>
</div>
</div>
<button id="LoginDialogBtn" type="button" data-toggle="modal" data-target="#LoginDialog" data-backdrop="static" style="visibility: hidden; display: none;"></button>
</div>
</div>
<style>
        #BodyContent_RblSelectOptions.md-radio label {
            margin-bottom: 0px;
        }

        h6 > span {
            color: red;
            font-size: 1rem;
            text-align: justify;
        }

        .styleBox h6 {
            font-weight: 400;
            margin-bottom: 25px;
            text-align: justify;
        }
        .container.templogin .temploginlogo {
    text-align: center;
    margin: 0 auto;
    padding: 12px 0 10px 0;
}
    </style>
</div>
</form>
</body>
</html>
`

var (
	district string
	interval int64
)

func main() {
	stopCh := signals.SetupSignalHandler()
	ticker := time.NewTicker(time.Duration(interval) * time.Minute)
	flag.Parse()
	go func() {
		for {
			select {
			case <-stopCh:
				return
			case <-ticker.C:
				findDistrictAndNotify(district)
			}
		}
	}()
	<-stopCh
}

func findDistrictAndNotify(district string) {
	resp, err := http.Get("https://int.keellssuper.net/login")
	if err != nil {
		log.Printf("Cannot call to keellssuper %v\n", err)
		return
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("Cannot parse html %v\n", err)
		return
	}
	found := false
	doc.Find("#BodyContent_ddlDeliveryCity").Each(func(idxSel int, selection *goquery.Selection) {
		selection.Find("option").Each(func(idxOpt int, opt *goquery.Selection) {
			if strings.ToLower(opt.Text()) == strings.ToLower(district) {
				found = true
				log.Printf("Found district %s\n", district)
			}
		})
	})
	if found {
		err = beeep.Notify("keells-check", fmt.Sprintf("%q district available", district), "")
		if err != nil {
			log.Printf("Cannot send desktop notification %v\n", err)
			return
		}
	}
}

func init() {
	flag.StringVar(&district, "d", "Matara", "District name to check")
	flag.Int64Var(&interval, "i", 1, "Check interval in minutes")
}

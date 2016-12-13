<?php
/* 
* 
* This script send email when user fills in the form
* Check out this tutorial on PHP forms http://myphpform.com/php-form-tutorial.php or this http://www.tizag.com/phpT/forms.php one
*
*/

$to = "office@example.com";
$subject = "Signup from Ventcamp";
$newsletter = array(true => 'Yes!', 0 => 'No please');
$message = "Ticket type: " . $_POST['ticket'];
$message .= "<br/>Fullname: " . $_POST['fullname'];
$message .= "<br>Company: " . $_POST['company'];
$message .= "<br>Email: " . $_POST['email'];
$message .= "<br>Newsletter subscription: " . $newsletter[$_POST['newsletter']];

$headers  = "MIME-Version: 1.0" . "\r\n";
$headers .= "Content-type: text/html; charset=utf-8" . "\r\n";
$headers .= "From: " . $_POST['fullname'] . " <" . $_POST['email'] . ">". "\r\n";

if( mail($to, $subject, $message, $headers) ) {
	echo "ok";
} else {
	echo "error";
}
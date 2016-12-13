<?php
/*
*
* Mailchimp subscription processing script
*
*/

// grab an API Key from http://admin.mailchimp.com/account/api/
$api_key = 'your-api-code-here';

// grab your List's Unique Id by going to http://admin.mailchimp.com/lists/
// Click the "settings" link for the list - the Unique Id is at the bottom of that page.
$list_id = "your-list-id-here";


//Making API call, no change needed here
if ( !empty($_GET['ajax']) ) {

	// Validation
	if (empty($_GET['email'])) {
		echo "No email address provided";
	}

	if(!preg_match("/^[_a-z0-9-]+(\.[_a-z0-9-]+)*@[a-z0-9-]+(\.[a-z0-9-]+)*$/i", $_GET['email'])) {
		echo "Email address is invalid";
	}

	require_once('assets/inc/mailchimp/MailChimp.php');

	$api = new MailChimp($api_key);

	$merge_fields = array(
		'FNAME' => !empty($_GET['fname']) ? $_GET['fname'] : '',
		'LNAME' => !empty($_GET['lname']) ? $_GET['lname'] : '',
	);

	$result = $api->post("lists/{$list_id}/members", array(
		'email_address'     => $_GET['email'],
		'status'            => 'subscribed',
		'merge_fields'      => $merge_fields
	), 100);

	if (!empty($result['status']) && $result['status'] === 'subscribed') {
		echo 'Success! Check your email to confirm sign up.';
	} else {
		$message = array();

		if( !empty($result['title']) ) {
			$message[] = $result['title'];
		}

		if( !empty($result['detail']) ) {
			$message[] = $result['detail'];
		}

		echo 'Error: '.implode(' - ', $message);
	}

	
	// After Mailchimp API call executed successfully you can send more emails
	// to yourself or users. Just uncomment and edit the code below
	/*
	$to = "office@example.com";
	$subject = "test email";
	$message .= "thanks for subscribing to our newsletter!"
	$headers  = "MIME-Version: 1.0" . "\r\n";
	$headers .= "Content-type: text/html; charset=utf-8" . "\r\n";
	$headers .= "From: Ventcamp Forms <" . $_GET['email'] . ">". "\r\n";

	if( mail($to, $subject, $message, $headers) ) {
		echo "ok";
	} else {
		echo "error";
	}*/

}
?>

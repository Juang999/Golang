<?php

$address1 = [
    "City" => "Bogor",
    "Province" => "East Java",
    "Country" => "Indonesia"
];
$address2 = $address1;

$address2['City'] = "Jakarta";

print_r($address1);
print_r($address2);
#!/bin/bash

# Variables
TO=(" " " ")
SUBJECT="Generated report for component performance tests"
BODY="Please find attached the report for component performance tests."
ATTACHMENT="file.txt"
FROM=""

#!/bin/bash

# Variables
SMTP_SERVER=""
SMTP_USER="your-email@example.com"
SMTP_PASSWORD="your-email-password"
# TO="recipient@example.com"
SUBJECT="Subject of the Email"
BODY="This is the body of the email."
ATTACHMENT="file.txt"
# FROM="your-email@example.com"

# Check if the attachment exists
if [ ! -f "$ATTACHMENT" ]; then
  echo "Attachment file not found: $ATTACHMENT"
  exit 1
fi

# Create email headers and body in a temporary file
email_content=$(mktemp)

echo "From: $FROM" > $email_content
echo "To: $TO" >> $email_content
echo "Subject: $SUBJECT" >> $email_content
echo "MIME-Version: 1.0" >> $email_content
echo "Content-Type: multipart/mixed; boundary=\"frontier\"" >> $email_content
echo "" >> $email_content
echo "--frontier" >> $email_content
echo "Content-Type: text/plain" >> $email_content
echo "" >> $email_content
echo "$BODY" >> $email_content
echo "" >> $email_content
echo "--frontier" >> $email_content
echo "Content-Type: text/plain; name=\"$(basename "$ATTACHMENT")\"" >> $email_content
echo "Content-Disposition: attachment; filename=\"$(basename "$ATTACHMENT")\"" >> $email_content
echo "Content-Transfer-Encoding: base64" >> $email_content
echo "" >> $email_content
base64 "$ATTACHMENT" >> $email_content
echo "" >> $email_content
echo "--frontier--" >> $email_content

# Send email using curl
curl -k --url "$SMTP_SERVER" \
     --mail-from "$FROM" \
     $(for RECIPIENT in "${TO_RECIPIENTS[@]}"; do echo "--mail-rcpt $RECIPIENT "; done) \
     --upload-file $email_content \
     --verbose
    #  --trace-ascii curl_trace.log
    #  --ssl-reqd
    #  --user "$SMTP_USER:$SMTP_PASSWORD"

# Clean up the temporary email file
# rm -f $email_content


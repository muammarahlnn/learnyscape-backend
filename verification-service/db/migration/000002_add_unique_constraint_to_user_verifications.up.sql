ALTER TABLE "user_verifications"
ADD CONSTRAINT "ux_user_verifications_user_id" UNIQUE ("user_id");
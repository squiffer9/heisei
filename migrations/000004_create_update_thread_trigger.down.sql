CREATE OR REPLACE FUNCTION update_thread_on_post() RETURNS TRIGGER AS $$
BEGIN
  UPDATE threads
  SET last_post_at = NEW.created_at,
      post_count = post_count + 1,
      updated_at = CURRENT_TIMESTAMP
  WHERE id = NEW.thread_id;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_thread_on_post
AFTER INSERT ON posts
FOR EACH ROW
EXECUTE FUNCTION update_thread_on_post();

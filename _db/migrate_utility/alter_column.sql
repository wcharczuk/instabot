DROP FUNCTION IF EXISTS alter_column(text,text,text);
CREATE FUNCTION alter_column(input_table_name text, input_column_name text, statement text) RETURNS void AS
$_$
BEGIN

IF EXISTS (
		SELECT 	1
		FROM	information_schema.columns c
		WHERE	c.table_name  = input_table_name and c.column_name = input_column_name
	) THEN
	EXECUTE statement;
	RAISE NOTICE 'MIGRATE :: ALTERED COLUMN % ON TABLE %', input_column_name, input_table_name;
ELSE
	RAISE EXCEPTION 'MIGRATE :: COLUMN % ON TABLE % DOESNT EXIST', input_column_name, input_table_name;
END IF;

END;
$_$ LANGUAGE plpgsql;
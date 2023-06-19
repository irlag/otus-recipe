CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE recipe (
      id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
      name varchar(255) NOT NULL,
      description text,
      cooking_time smallint NOT NULL,
      calories int,
      proteins int,
      fats int,
      carbohydrates int,
      rating numeric(2,1) NOT NULL DEFAULT '0',
      vote_count integer NOT NULL DEFAULT '0',
      vote_sum integer NOT NULL DEFAULT '0',
      version uuid NOT NULL DEFAULT gen_random_uuid()
);

COMMENT ON COLUMN public.recipe.cooking_time IS 'Time for cooking in minutes';
COMMENT ON COLUMN public.recipe.calories IS 'Dish calories in kcal';
COMMENT ON COLUMN public.recipe.proteins IS 'Dish fats in grams';
COMMENT ON COLUMN public.recipe.fats IS 'Dish fats in grams';
COMMENT ON COLUMN public.recipe.carbohydrates IS 'Dish carbohydrates in grams';

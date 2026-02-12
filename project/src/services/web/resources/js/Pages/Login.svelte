<script lang="js">
  import { Button } from "$lib/components/ui/button";
  import {
    Card,
    CardContent,
    CardDescription,
    CardHeader,
    CardTitle,
  } from "$lib/components/ui/card";
  import * as Field from "$lib/components/ui/field/index.js";
  import { Input } from "$lib/components/ui/input/index.js";
  import { Spinner } from "$lib/components/ui/spinner/index.js";
  import { useForm } from "@inertiajs/svelte";

  import * as yup from "yup";

  let loading = $state(false);

  // track touched fields
  let touched = $state({
    email: false,
    password: false,
  });

  // Yup schema
  const schema = yup.object({
    email: yup.string().email("Email inválido").required("Email é obrigatório"),
    password: yup
      .string()
      .min(4, "A senha deve ter no mínimo 4 caracteres")
      .required("Senha é obrigatória"),
  });

  const form = useForm({
    email: null,
    password: null,
  });

  async function validateField(field) {
    try {
      await schema.validateAt(field, $form);
      $form.clearErrors(field);
    } catch (err) {
      $form.setError(field, err.message);
    }
  }

  async function handleBlur(field) {
    touched[field] = true;
    await validateField(field);
  }

  async function handleInput(field) {
    if (touched[field]) {
      await validateField(field);
    }
  }

  async function handleSubmit(e) {
    e.preventDefault();
    loading = true;
    await setTimeout(function () {
      schema
        .validate($form, { abortEarly: false })
        .then(() => {
          loading = false;
          console.log($form.email);
          console.log($form.password);
          $form.post("/login")
        })
        .catch((err) => {
          loading = false;
          err.inner.forEach((e) => {
            $form.setError(e.path, e.message);
          });
        });
    }, 1000);
  }
</script>

<div class="min-h-screen flex items-center justify-center bg-muted px-4">
  <Card class="w-full max-w-md shadow-lg">
    <CardHeader class="space-y-1">
      <CardTitle class="text-2xl text-center">Login</CardTitle>
      <CardDescription class="text-center">
        Entre com o seu email e senha
      </CardDescription>
    </CardHeader>

    <CardContent>
      <form onsubmit={handleSubmit} class="space-y-4">
        <Field.Set>
          <Field.Group>
            <Field.Field>
              <Field.Label for="email">Email</Field.Label>
              <Input
                id="email"
                type="text"
                onblur={() => handleBlur("email")}
                oninput={() => handleInput("email")}
                bind:value={$form.email}
                placeholder="meu@email.com"
              />
              <!-- <Field.Description
                >Entre com o seu email</Field.Description
              > -->
              {#if $form.errors.email}
                <Field.Error>
                  {$form.errors.email}
                </Field.Error>
              {/if}
            </Field.Field>

            <Field.Field>
              <Field.Label for="password">Senha</Field.Label>
              <!-- <Field.Description
                >Deve ter no mínimo 4 caracteres.</Field.Description
              > -->
              <Input
                id="password"
                type="password"
                bind:value={$form.password}
                onblur={() => handleBlur("password")}
                oninput={() => handleInput("password")}
                placeholder="••••••••"
              />
              {#if $form.errors.password}
                <Field.Error>
                  {$form.errors.password}
                </Field.Error>
              {/if}
            </Field.Field>

            <Field.Field orientation="responsive">
              <Button class="w-full" disabled={loading} type="submit">
                {#if loading}
                  <Spinner />
                  Enviando
                {:else}
                  Login
                {/if}
              </Button>
            </Field.Field>
          </Field.Group>
        </Field.Set>
      </form>
    </CardContent>
  </Card>
</div>

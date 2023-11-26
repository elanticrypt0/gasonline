#!/bin/bash

mkdir ./wui
cd wui

echo 'Astro install'


npm create astro@latest && npx astro add tailwind && npx astro add svelte && npm install flowbite
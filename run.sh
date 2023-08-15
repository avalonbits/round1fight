#!/bin/sh
export $(cat .env | xargs) && ./round1fight

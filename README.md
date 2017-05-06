# Amazon Alexa + IRKit in Golang

Hello,

This is a sample implementation of a custom [Alexa Smart Home Skill](https://developer.amazon.com/public/solutions/alexa/alexa-skills-kit/docs/smart-home-skill-api-reference) using Golang. Alexa recognizes your voice commands and controls your IR-enabled appliances (in this case, Roomba, AC, and lights) by sending an infrared signal via [IRKit](http://getirkit.com/#IRKit-Internet-API). 

The code was extended slightly to handle multiple appliances and more commands from [Taichi Nakashima](https://github.com/tcnksm)’s [original implementation](https://github.com/tcnksm/alexa-irkit-ac)

See demo on [https://www.youtube.com/watch?v=-4y4qR_PkOM](https://www.youtube.com/watch?v=-4y4qR_PkOM).

## Deploy

To deploy the custom skill (AWS labmda function), use [apex](https://github.com/apex/apex):

```bash
$ apex deploy ac
```

## Author 
[Vincent Ethier](https://github.com/vinidlidoo)





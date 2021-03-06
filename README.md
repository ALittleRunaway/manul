# Manul

Hi! Meet manul - an ASCII-image converter. \
Convert any image you like to it's ASCII representation.

## Installation 
```bash
git clone https://github.com/ALittleRunaway/manul.git
cd manul
go install manul # make sure your $GOPATH variable is in $PATH
```

## Instructions

```bash
Manul is a CLI tool for converting your images to ASCII-symbols.
Supported image formats: PNG, JPG, GPEG.

Usage:
  manul [options...] <filename> [flags]

Examples:
    manul myimage.png
    manul -w 50 -o ~/myimageascii.txt myimage.png
    manul -i -n 100 myimage.png

Flags:
    -n, --height int          Height in pixels. Default is 200. If both width and height are provided, height is ignored
    -h, --help                help for manul
    -i, --invert              Invert image colors
    -o, --outputFile string   Output file name. Default is 'ascii.txt' (default "ascii.txt")
    -v, --version             version for manul
    -w, --width int           Width in pixels. Default is 300. If both width and height are provided, height is ignored
```

### Some examples

|                                                              Original Joker                                                               |                                                                ASCII-Joker                                                                 |
|:-----------------------------------------------------------------------------------------------------------------------------------------:|:------------------------------------------------------------------------------------------------------------------------------------------:|
| <img width="760" alt="image" src="https://user-images.githubusercontent.com/66780462/156937158-65d75a12-3fcb-4ad1-b7e1-b55febe01f35.png"> | <img width="1470" alt="image" src="https://user-images.githubusercontent.com/66780462/156936519-4c5bbdfc-df3c-482c-a2d4-2a70f832bc9c.png"> |
|                                                              Original Manul                                                               |                                                                ASCII-Manul                                                                 |
|       <img alt="image" src="https://user-images.githubusercontent.com/66780462/156936584-01783915-9982-44b8-99b5-718b19810936.png">       | <img width="1470" alt="image" src="https://user-images.githubusercontent.com/66780462/156936546-d880b444-90ac-4d43-9579-c0777adf640d.png"> |
|                                                         Original Legendary Album                                                          |                                                           ASCII-Legendary Album                                                            |
|       <img alt="image" src="https://user-images.githubusercontent.com/66780462/156936564-84647d59-6d4b-4e6e-8ae4-a8d864984c93.png">       | <img width="1470" alt="image" src="https://user-images.githubusercontent.com/66780462/156936531-d45988e9-80d4-491e-94ad-ed733439a052.png"> |

### An ASCII-Manul

```
,,,,::::::::;;;;;;;;::::,,,,,,::;;::::;;::,,&&&&xxXXXX::::XXxx::::xx&&xx;;::;;;;xxxx;;;;;;;;;;;;;;::
::::::::::::;;::;;;;::,,,,,,..,,::::;;::::::@@XXxxXXXX;;;;::,,::;;;;,,::::::::::;;;;;;;;;;::::;;::,,
;;::::::::::::;;;;::,,::xxxx..,,,,,,::::;;xxXX::xx;;;;xxxx::::::::,,,,::;;;;;;::;;;;;;;;::::;;;;;;xx
..,,,,,,::,,::;;::,,xxXX;;xx;;..,,::::::::;;;;::;;;;::,,::::,,,,,,..xxxxXXXXXX;;::;;xxxx;;::::::::::
,,,,,,,,,,,,::;;::::;;::,,,,::,,,,::..,,..,,::;;;;::;;,,,,,,..,,,,,,::::::::;;::,,::;;;;;;;;::,,,,,,
,,,,,,::,,::::;;;;::::,,,,,,,,,,,,::,,,,,,,,::xx&&;;;;::::,,,,;;::::;;::,,,,,,,,::;;;;;;xxxx;;::,,,,
::;;::::,,::::;;;;::,,::xxXXXXXX;;::;;xx;;;;;;xx&&xx;;;;xxxx::;;xx&&&&XXXXxx::::;;XXxx;;;;xx;;::::,,
;;::::,,::;;xx;;::::::&&####&&XX&&XX;;&&&&xx;;xx&&XXxx&&@@XX;;XX@@@@&&&&@@##XX;;;;XX;;;;;;;;xx;;::::
,,..,,::::;;;;;;xxxxXX@@XXXX;;XX##@@xx;;XXXXxxXX@@xxXX&&XX;;XX##&&XX@@XXxx@@##xxxxxxxx;;;;xxXXxx;;::
,,,,,,::;;xxxxxxxxXXXXXXxx;;::&&##XX&&xx;;XX&&&&&&XXXXXX::xx##@@;;xx##XX;;&&&&XXXXxxxxxx;;;;xxxx;;;;
,,,,,,::xxxx;;xxxxXX::;;XXxx::::;;xx@@&&::;;XXXX&&XXxx;;::XX##@@;;::::::::&&;;;;XX&&XX;;xxXXxxXXxxxx
,,,,::::;;::;;;;xx;;,,..;;&&xxxxXX&&##@@::;;xxxxXXxxxx::::&&####&&xx;;xxXX::..::::;;&&XX;;;;xxxxXXxx
::,,,,::;;;;;;XXxx::,,,,....;;XX&&XXxxxx;;;;;;;;xx;;::::::xxxxXX&&&&xx::,,,,;;;;::::;;XXxx;;xx;;xxxx
::::::xxxxxxXXxx,,,,::::;;::..  ......;;;;;;,,,,::,,,,::;;;;,,......,,::::::::::..,,::xx&&XXXXXXXXxx
::::::;;;;XXXX::,,....,,,,::;;;;::..,,;;;;::::,,::,,,,::xx;;,,..,,::;;::::,,,,,,,,::::::;;XX@@@@@@&&
,,,,::xxxx;;,,..,,,,::::::;;xx&&xx,,;;xx;;::::::::::,,::;;xx;;,,xx&&xxxx;;;;;;::,,,,,,,,,,;;;;XX&&&&
,,,,;;;;;;,,....,,;;&&&&xx;;;;;;::::xx;;::::::::::::::::;;xx;;;;::xx&&&&XX@@@@xx;;,,,,..,,,,::::;;;;
::;;;;;;::,,,,,,,,;;xxxxXXXX;;::::;;xxXX::::;;xxxx;;::;;xxXX;;;;;;::xxXX&&XXxxXXxx;;,,,,....,,,,..,,
;;;;;;;;,,,,,,,,::xxxx::,,::::::;;::;;;;::;;::;;;;::::::::::::;;xx::::;;;;xxXXxx;;::,,,,,,,,,,,,,,..
;;xx;;,,....,,::;;;;;;xx::,,::;;xx::,,,,::;;::;;;;::;;;;,,,,,,::XXxx::,,,,::;;::::;;xx;;xxxx::::::,,
;;::,,....,,::xxXX;;::::::,,::xxxx,,,,::;;xx;;;;;;XXxxxx;;,,,,,,xxXXxx::::::;;;;xxxxXXXX&&&&XXxx;;::
;;;;;;;;;;xxXXXX;;;;;;::::::xx;;::;;;;;;XX####XX&&####&&;;;;::;;::;;XXxx::;;;;;;;;;;;;;;;;XX&&@@&&XX
;;xx;;;;;;XXXXxx;;;;;;::::::::::::::::::::XX@@&&&&@@xx,,,,::::,,,,,,..,,::::::::::::::::::;;XXXXXXXX
xxxx,,,,::xxxx;;;;::::::,,,,,,::::,,::,,,,::XX@@@@xx::::::,,,,,,::::::,,,,::::::::::::::::::;;xxXXxx
XX;;,,::xxxx;;::::::::,,,,::::;;;;;;;;;;::::;;xx;;;;;;::;;::::;;::,,::,,,,,,::,,,,::::;;;;::::::::;;
;;::::;;;;::::::;;;;::::,,::::::::::;;xx;;;;;;;;;;xx::;;;;;;;;::::::::,,,,::;;::::::::::::;;::::::;;
::,,,,;;;;;;;;xx;;::::::::::::::;;;;;;;;;;;;XX&&@@&&XXxx;;;;;;;;::::::,,::;;;;;;;;;;::::::::::;;;;;;
,,::;;xxxxXXXX;;::;;;;;;xx;;;;::;;xxxxxxXX@@########@@@@XX;;;;xxxx;;;;::::::;;::::;;;;;;::::,,,,::::
::;;;;XX&&XX::::xx;;xxXXxx;;;;xxxxxxxx&&@@@@&&XXXXXX&&@@@@&&xxxxxx;;;;;;;;;;::::::::::::::,,,,,,,,::
;;;;xxXXxx::xxxxxxxxxx;;;;xxxxxx;;xxxx&&&&XXXXXXxxxxxxxxXX&&XXxxxxxx;;;;;;;;;;::::,,,,,,,,,,,,,,::;;
xx;;xx;;;;&&XXxxxx;;;;;;;;;;;;;;;;;;;;xx;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;::::::,,,,,,,,,,,,,,,,::
;;;;;;xx&&XXxxxx;;;;;;xxxx;;;;::;;;;;;;;;;;;;;;;;;;;;;;;;;;;::::::;;;;;;;;;;::::::::,,,,,,,,,,::::::
```

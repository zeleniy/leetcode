<?php

namespace Zeleniy\Leetcode\ValidAnagram;

class DataSet
{
    public static function getData() {

        return [
            [true, 'anagram', 'nagaram'],
            [true, 'hydroxydeoxycorticosterones', 'hydroxydesoxycorticosterone'],
            [false, 'rat', 'car'],
            [false, 'aacc', 'ccac'],
            [
                false,
                'thecomplexofpalaceswasrebuiltandexpandedseveraltimesduringitshistorymuchofthecomplexwasdestroyedduringthenikariotsofandwasrebuiltlavishlybytheemperorjustinianifurtherextensionsandalterationswerecommissionedbyjustinianiiandbasilihoweverithadfallenintodisrepairbythetimeofconstantineviiwhoordereditsrenovationfromtheearlycenturyonwardsemperorsfavouredthepalaceofblachernaeasanimperialresidencethoughtheycontinuedtousethegreatpalaceastheprimaryadministrativeandceremonialcentreofthecityitdeclinedsubstantiallyduringthefollowingcenturywhenpartsofthecomplexweredemolishedorfilledwithrubbleduringthesackofconstantinoplebythefourthcrusadethepalacewasplunderedbythesoldiersofbonifaceofmontferratalthoughthesubsequentlatinemperorscontinuedtousethepalacecomplextheylackedmoneyforitsmaintenancethelastlatinemperorbaldwiniiwentasfarasremovingtheleadroofsofthepalaceandsellingthem',
                'thecomplexofpalaceswasrebuiltandexpandedseveraltimesduringitshistorymuchofthecomplexwasdestroyedduringthenikariotsofandwasrebuiltlavishlybytheemperorjustinianifurtherextensionsandalterationswerecommissionedbyjustinianiiandbasilihoweverithadfallenintodisrepairbythetimeofconstantineviiwhoordereditsrenovationfromtheearlycenturyonwardsemperorsfavouredthepalaceofblachernaeasanimperialresidencethoughtheycontinuedtousethegreatpalaceastheprimaryadministrativeandceremonialcentreofthecityitdeclinedsubstantiallyduringthefollowingcenturywhenpartsofthecomplexweredemolishedorfilledwithrubbleduringthesackofconstantinoplebythefourthcrusadethepalacewasplunderedbythesoldiersofbonifaceofmontferratalthoughthesubsequentlatinemperorscontinuedtousethepalacecomplextheylackedmoneyforitsmaintenancethelastlatinemperorbaldwiniiwentasfarasremovingtheleadroofsofthepalaceandsellingthen',
            ]
        ];
    }
}
